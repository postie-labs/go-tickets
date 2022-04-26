package qr

import (
	"fmt"
	"image"
	"os"

	"github.com/postie-labs/proto/qr"
)

func Read(filename string) (*qr.Code, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", img)

	return nil, nil
}
