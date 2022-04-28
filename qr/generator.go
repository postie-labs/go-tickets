package qr

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/postie-labs/proto/qr"
	"github.com/skip2/go-qrcode"
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
	dataBytes, err := proto.Marshal(code)
	if err != nil {
		return err
	}
	dataStr := base64.StdEncoding.EncodeToString(dataBytes)
	qc, err := qrcode.New(dataStr, qrcode.Medium)
	if err != nil {
		return err
	}
	fmt.Println(dataStr)
	err = qc.WriteFile(256, filename)
	if err != nil {
		return err
	}
	return nil
}
