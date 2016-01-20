package foswiki

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Base64ImageWriter interface {
	WriteImage(string, *base64Image) error
}

type base64ImageWriter struct{}

func NewBase64ImageWriter() *base64ImageWriter {
	return &base64ImageWriter{}
}

func (i *base64ImageWriter) WriteImage(directory string, image *base64Image) error {
	// create directory if it doesnt exist
	if _, err := os.Stat(directory); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(directory, 0755); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// write image to disk
	imageFilePath := fmt.Sprintf("%v/%v", directory, image.fileName)
	fmt.Println("creating image", imageFilePath)
	err := ioutil.WriteFile(filepath.Clean(imageFilePath), image.data, 0644)

	return err
}
