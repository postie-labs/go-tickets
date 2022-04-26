package qr

import (
	"fmt"

	"github.com/postie-labs/proto/qr"
	"github.com/skip2/go-qrcode"
	"github.com/tendermint/tendermint/crypto"
	"google.golang.org/protobuf/proto"
)

func Generate(owner, tokenId string, privKey crypto.PrivKey) (*qr.Code, error) {
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
	codeStr := code.String()
	qc, err := qrcode.New(codeStr, qrcode.Medium)
	if err != nil {
		return err
	}
	fmt.Printf("len:%d\ndata:%s\n%s", len(codeStr), codeStr, qc.ToSmallString(false))
	err = qc.WriteFile(256, filename)
	if err != nil {
		return err
	}
	return nil
}
