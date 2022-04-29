package main

import "github.com/postie-labs/go-tickets/qr"

const (
	DefaultQRCodeFilename = "qr-code.png"
)

func main() {
	code, err := qr.Read(DefaultQRCodeFilename)
	if err != nil {
		panic(err)
	}

	_, err = qr.Scan(code)
	if err != nil {
		panic(err)
	}
}
