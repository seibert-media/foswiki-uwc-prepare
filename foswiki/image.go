package foswiki

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type ImageConverter struct {
	Document   Doc
	imageIndex int
	images     []fImage
}

type fImage struct {
	fileName string
	size     int
	width    int
	height   int
}

func (c *ImageConverter) ReplaceBase64Tag(imageTag []byte) []byte {
	var imageTagReg = regexp.MustCompile(`(?si)<img .*?src="data:image/(.+?);base64,(.+?)".*?/>`)
	var err error

	matchedParts := imageTagReg.FindSubmatch(imageTag)
	if matchedParts == nil {
		return imageTag
	}

	image := c.addNewImage()
	image.fileName = fmt.Sprintf("base64-image%04d.%v", c.imageIndex, string(matchedParts[1]))

	image.width, image.height, image.size, err = c.saveBase64Image(matchedParts[2])
	if err != nil {
		panic(err)
	}

	imageSizeAttributes := ""
	if image.width > 0 && image.height > 0 {
		imageSizeAttributes = fmt.Sprintf(`width="%v" height="%v" `, image.width, image.height)
	}

	replacedTag := fmt.Sprintf(`<img alt="" src="%%ATTACHURLPATH%%/%v/%v/%v" %v/>`,
		c.Document.WebName(), c.Document.PageName(), image.fileName, imageSizeAttributes)
	c.imageIndex++

	return []byte(replacedTag)
}

func (c *ImageConverter) MetaData() []byte {
	metaData := ""
	date := 1451606400 // 01.01.2016
	username := "admin"
	version := 1
	for _, image := range c.images {
		metaData += fmt.Sprintf(`%%META:FILEATTACHMENT{name="%v" attachment="%v" attr="" comment="" date="%v" path="%v" size="%v" stream="%v" user="%v" version="%v"}%%
`, image.fileName, image.fileName, date, image.fileName, image.size, image.fileName, username, version)
	}
	return []byte(metaData)
}

func (c *ImageConverter) addNewImage() *fImage {
	c.images = append(c.images, fImage{})
	return &c.images[c.imageIndex]
}

func (i *ImageConverter) saveBase64Image(base64Data []byte) (w int, h int, size int, err error) {

	imageData, err := base64.StdEncoding.DecodeString(string(base64Data))
	if err != nil {
		return
	}
	size = len(imageData)

	imageReader := bytes.NewReader(imageData)
	image, _, err := image.DecodeConfig(imageReader)
	if err == nil {
		w, h = image.Width, image.Height
	}

	if _, err = os.Stat(i.Document.PubDir()); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(i.Document.PubDir(), 0755); err != nil {
				return
			}
		} else {
			return
		}
	}

	imageFilePath := fmt.Sprintf("%v/%v", i.Document.PubDir(), i.images[i.imageIndex].fileName)
	fmt.Println("creating image", imageFilePath)
	err = ioutil.WriteFile(filepath.Clean(imageFilePath), imageData, 0644)

	return
}
