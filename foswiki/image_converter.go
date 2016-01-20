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
	Document Doc
	images   []Image
}

// ReplaceBase64Tag replace HTML image tags with base64-embedded pictures.
// The images are written to the corresponding Foswiki pub folders.
// It returns the replaced HTML image tag as byte array.
func (c *ImageConverter) ReplaceBase64Tag(imageTag []byte) []byte {
	// extract information from the HTML image tag
	var re = regexp.MustCompile(`(?si)<img .*?src="data:image/(.+?);base64,(.+?)".*?/>`)
	matchedParts := re.FindSubmatch(imageTag)
	if matchedParts == nil {
		return imageTag
	}
	fileExtension := string(matchedParts[1])
	base64Data := matchedParts[2]

	// decode image and write to disk
	err := c.saveBase64Image(base64Data, fileExtension)
	if err != nil {
		panic(err)
	}

	// return replaced HTML image tag
	replacedTag := fmt.Sprintf(`<img alt="" src="%%ATTACHURLPATH%%/%v/%v/%v" %v/>`,
		c.Document.WebName(), c.Document.PageName(), c.lastImage().fileName, c.lastImage().DimensionHTMLAttributes())

	return []byte(replacedTag)
}

// MetaData returns the Foswiki metadata HTML code for all decoded base64-embedded images.
func (c *ImageConverter) MetaDataHTML() []byte {
	metaData := ""
	date := 1451606400 // 01.01.2016
	username := "admin"
	version := 1

	for _, image := range c.images {
		metaData += fmt.Sprintf(
			`%%META:FILEATTACHMENT{name="%v" attachment="%v" attr="" comment="" date="%v" path="%v" size="%v" stream="%v" user="%v" version="%v"}%%%v`,
			image.fileName, image.fileName, date, image.fileName, image.size, image.fileName, username, version, "\n")
	}
	return []byte(metaData)
}

// newImage adds a new fImage struct to ImageConverter images array.
// It set a filename with a counter and returns a pointer to the new fImage struct.
func (c *ImageConverter) newImage(fileExtension string) *Image {
	c.images = append(c.images, Image{})
	newImage := c.lastImage()
	newImage.fileName = fmt.Sprintf("base64-image%04d.%v", len(c.images)-1, fileExtension)
	return newImage
}

// lastImage returns the last fImage struct of the ImageConverter images array.
func (c *ImageConverter) lastImage() *Image {
	return &c.images[len(c.images)-1]
}

// saveBase64Image decodes a base64 byte array to a image file in the corresponding Foswiki pub folder.
// It saves the image metadata in a new added fImage struct.
func (c *ImageConverter) saveBase64Image(base64Data []byte, fileExtension string) error {
	imageData, err := base64.StdEncoding.DecodeString(string(base64Data))
	if err != nil {
		return err
	}

	// determine and save image metadata
	newImage := c.newImage(fileExtension)
	newImage.size = len(imageData)
	imageReader := bytes.NewReader(imageData)
	imageConfig, _, err := image.DecodeConfig(imageReader)
	if err == nil {
		newImage.width, newImage.height = imageConfig.Width, imageConfig.Height
	}

	// create missing pub folder structure
	if _, err = os.Stat(c.Document.PubDir()); err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(c.Document.PubDir(), 0755); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// write image to disk
	imageFilePath := fmt.Sprintf("%v/%v", c.Document.PubDir(), newImage.fileName)
	fmt.Println("creating image", imageFilePath)
	err = ioutil.WriteFile(filepath.Clean(imageFilePath), imageData, 0644)

	return err
}
