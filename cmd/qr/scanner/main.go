package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/postie-labs/go-tickets/qr"
)

const (
	DefaultQRCodeFilename = "qr-code.png"
)

func main() {
	code, err := qr.Read(DefaultQRCodeFilename)
	if err != nil {
		panic(err)
	}

	ok, err := qr.Scan(code)
	if err != nil {
		panic(err)
	}

	fmt.Println(ok)
}

func init() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount("terra", "terrapub")
	config.Seal()
}
