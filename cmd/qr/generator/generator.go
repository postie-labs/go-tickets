package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/postie-labs/proto/qr"
	"github.com/tendermint/tendermint/crypto"
)

func GenerateQRCode(owner, tokenId string, privKey crypto.PrivKey) (*qr.Code, error) {
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
		PublicKey: string(privKey.PubKey().Bytes()),
	}
	return &qr.Code{
		Data:      data,
		Signature: signature,
	}, nil
}

func PrintQRCode(code *qr.Code) {

}
