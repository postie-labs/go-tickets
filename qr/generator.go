package qr

import (
	"encoding/base64"
	"fmt"
	"image/png"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/postie-labs/proto/qr"
	"google.golang.org/protobuf/proto"
)

func Generate(owner, tokenId string, privKey types.PrivKey) (*qr.Code, error) {
	data := &qr.Data{
		Owner:   owner,
		TokenId: tokenId,
	}
	dataBytes, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}
	sigBytes, err := privKey.Sign(dataBytes)
	if err != nil {
		return nil, err
	}
	signature := &qr.Signature{
		SigBytes:  sigBytes,
		PublicKey: privKey.PubKey().Bytes(),
	}
	return &qr.Code{
		Data:      data,
		Signature: signature,
	}, nil
}

func Write(code *qr.Code, filename string) error {
	// marshal code to data bytes
	dataBytes, err := proto.Marshal(code)
	if err != nil {
		return err
	}

	// encode data bytes to base64-formatted data string
	dataStr := base64.StdEncoding.EncodeToString(dataBytes)
	fmt.Println(dataStr)

	// encode base64-formatted data string to qr code image
	writer := qrcode.NewQRCodeWriter()
	img, err := writer.EncodeWithoutHint(dataStr, gozxing.BarcodeFormat_QR_CODE, 512, 512)
	if err != nil {
		return err
	}

	// write qr code image to file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}
