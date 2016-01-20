package foswiki

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// base64Image stores decoded image bytes and image metadata
type base64Image struct {
	fileName string
	data     []byte
	size     int
	width    int
	height   int
}

// NewBase64 read a Foswiki image as base64 byte array and extract image metadata.
func NewBase64Image(fileName string, data []byte) (*base64Image, error) {
	var err error

	obj := base64Image{fileName: fileName}
	obj.data, err = base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	obj.size = len(obj.data)
	imageReader := bytes.NewReader(obj.data)
	imageConfig, _, err := image.DecodeConfig(imageReader)
	if err == nil {
		obj.width, obj.height = imageConfig.Width, imageConfig.Height
	}
	return &obj, nil
}

func (i *base64Image) FileName() string {
	return i.fileName
}

func (i *base64Image) Data() []byte {
	return i.data
}

func (i *base64Image) Size() int {
	return i.size
}

func (i *base64Image) Width() int {
	return i.width
}

func (i *base64Image) Height() int {
	return i.height
}

// dimensionHTMLAtrtibures returns the width and height attributes of a Foswiki Image as HTML code.
func (i *base64Image) DimensionHTMLAttributes() string {
	if i.width == 0 || i.height == 0 {
		return ""
	}
	return fmt.Sprintf(`width="%v" height="%v" `, i.width, i.height)
}
