package qr

import (
	"encoding/base64"
	"image"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/postie-labs/proto/qr"
	"google.golang.org/protobuf/proto"
)

func Read(filename string) (*qr.Code, error) {
	// read image file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decode file to qr code image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, err
	}

	// decode qr code image to base64-formatted data string
	reader := qrcode.NewQRCodeReader()
	result, err := reader.DecodeWithoutHints(bmp)
	if err != nil {
		return nil, err
	}

	// decode base64-formatted data string to data bytes
	dataStr := result.GetText()
	dataBytes, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		return nil, err
	}

	// unmarshal data bytes to qr code
	var code qr.Code
	err = proto.Unmarshal(dataBytes, &code)
	if err != nil {
		return nil, err
	}
	return &code, nil
}
