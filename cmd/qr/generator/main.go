package main

import (
	"github.com/terra-project/terra.go/key"
)

func main() {
	mnemonic, err := key.CreateMnemonic()
	if err != nil {
		panic(err)
	}
	privKeyBytes, err := key.DerivePrivKey(mnemonic, key.CreateHDPath(0, 0))
	if err != nil {
		panic(err)
	}
	privKey, err := key.StdPrivKeyGen(privKeyBytes)
	if err != nil {
		panic(err)
	}
	owner := "terra1k05lru8us3ctq7ngc396sxesmd7dsd2a8ppfv7"
	tokenId := "this is the first token issued by alice+bob"
	qrCode, err := GenerateQRCode(owner, tokenId, privKey)
	if err != nil {
		panic(err)
	}
	PrintQRCode(qrCode)
}
