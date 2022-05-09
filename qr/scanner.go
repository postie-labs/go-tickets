package qr

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/postie-labs/proto/qr"
	pb "github.com/postie-labs/proto/tickets"
	"google.golang.org/protobuf/proto"
)

const (
	DefaultLCDEndpoint = "https://bombay-lcd.terra.dev"
)

func Scan(code *qr.Code) (bool, error) {
	// 1. verify code.Data with code.Signature
	var pubKey secp256k1.PubKey
	err := pubKey.UnmarshalAmino(code.Signature.PublicKey)
	if err != nil {
		return false, err
	}
	dataBytes, err := proto.Marshal(code.Data)
	if err != nil {
		return false, err
	}
	if !pubKey.VerifySignature(dataBytes, code.Signature.SigBytes) {
		return false, fmt.Errorf("failed to verify signature")
	}
	if types.AccAddress(pubKey.Address()).String() != code.Data.Owner {
		return false, fmt.Errorf("failed to verify owner address")
	}

	// 2. get ticket metadata
	query := pb.QueryAllInfo{AllNftInfo: &pb.QueryAllInfo_AllInfo{TokenId: code.Data.TokenId}}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return false, err
	}
	queryStr := base64.StdEncoding.EncodeToString(queryBytes)
	urlPath := fmt.Sprintf("%s/%s/%s/%s?query_msg=%s",
		DefaultLCDEndpoint,
		"terra/wasm/v1beta1/contracts",
		code.Data.ContractAddress,
		"store",
		queryStr,
	)
	resp, err := http.Get(urlPath)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var queryResp pb.QueryAllInfoResponse
	err = json.Unmarshal(data, &queryResp)
	if err != nil {
		return false, err
	}

	// 3. check ownership
	if code.Data.Owner != queryResp.QueryResult.Access.Owner {
		return false, fmt.Errorf("failed to verify ownership")
	}

	// 4. check validity with not_valid_before, not_valid_after
	now := time.Now().Unix()
	notValidBefore := queryResp.QueryResult.Info.Extension.NotValidBefore
	notValidAfter := queryResp.QueryResult.Info.Extension.NotValidAfter
	if !(notValidBefore <= now && now <= notValidAfter) {
		return false, fmt.Errorf("failed to verify validity")
	}

	// TODO: 5. check attributes (optional)

	return true, nil
}

func Read(filename string) (*qr.Code, error) {
	// read image file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decode file to qr code image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, err
	}

	// decode qr code image to base64-formatted data string
	reader := qrcode.NewQRCodeReader()
	result, err := reader.DecodeWithoutHints(bmp)
	if err != nil {
		return nil, err
	}

	// decode base64-formatted data string to data bytes
	dataStr := result.GetText()
	dataBytes, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		return nil, err
	}

	// unmarshal data bytes to qr code
	var code qr.Code
	err = proto.Unmarshal(dataBytes, &code)
	if err != nil {
		return nil, err
	}
	return &code, nil
}
