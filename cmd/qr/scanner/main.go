package main

import "github.com/postie-labs/go-tickets/qr"

const (
	DefaultQRCodeFilename = "qr-code.png"
)

func main() {
	_, err := qr.Read(DefaultQRCodeFilename)
	if err != nil {
		panic(err)
	}
}
