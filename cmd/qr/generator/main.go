package main

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/postie-labs/go-tickets/qr"
	"github.com/terra-money/terra.go/key"
)

const (
	Mnemonic              = "this tiny rifle pelican board chalk adult voice imitate green daring visa grab amateur good finish yard note meat pink suit saddle catch alarm"
	ContractAddress       = "terra1al87aagg7asjyceu9x8f4xj554ddzlk9q2t8ls"
	Owner                 = "terra1jrj9kjwv5pwrttrsdmg33w0kzj3szzc230g9z5"
	TokenId               = "this is the first token issued by alice+bob"
	DefaultQRCodeFilename = "qr-code.png"
)

func main() {
	privKeyBytes, err := key.DerivePrivKeyBz(Mnemonic, key.CreateHDPath(0, 0))
	if err != nil {
		panic(err)
	}
	privKey, err := key.PrivKeyGen(privKeyBytes)
	if err != nil {
		panic(err)
	}
	qrCode, err := qr.Generate(ContractAddress, Owner, TokenId, privKey)
	if err != nil {
		panic(err)
	}
	err = qr.Write(qrCode, DefaultQRCodeFilename)
	if err != nil {
		panic(err)
	}
}

func init() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("terra", "terrapub")
	config.Seal()
}
