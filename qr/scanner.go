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
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return nil, err
	}
	reader := qrcode.NewQRCodeReader()
	result, err := reader.DecodeWithoutHints(bmp)
	if err != nil {
		return nil, err
	}

	dataStr := result.GetText()
	dataBytes, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		return nil, err
	}
	var data qr.Code
	err = proto.Unmarshal(dataBytes, &data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
