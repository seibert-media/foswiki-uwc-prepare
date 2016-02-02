package foswiki

import (
	"fmt"
	"regexp"
)

type imageConverter struct {
	foswikiDoc  Doc
	images      []base64Image
	imageWriter Base64ImageWriter
}

// NewImageConverter returns new imageConverter object.
func NewImageConverter(foswikiDoc Doc, imageWriter Base64ImageWriter) *imageConverter {
	return &imageConverter{foswikiDoc: foswikiDoc, imageWriter: imageWriter}
}

// ReplaceBase64Tag replace HTML image tags with base64-embedded pictures.
// The images are written to the corresponding Foswiki pub folders.
// It returns the replaced HTML image tag as byte array.
func (c *imageConverter) ReplaceBase64Tag(imageTag []byte) []byte {
	// extract information from the HTML image tag
	var re = regexp.MustCompile(`(?si)<img .*?src="data:image/(.+?);base64,(.+?)".*?/>`)
	matchedParts := re.FindSubmatch(imageTag)
	if matchedParts == nil {
		return imageTag
	}
	fileExtension := string(matchedParts[1])
	base64Data := matchedParts[2]

	// decode image and write to disk
	image, err := c.addImage(base64Data, fileExtension)
	if err != nil {
		panic(err)
	}

	// return replaced HTML image tag
	replacedTag := fmt.Sprintf(`<img alt="" src="%%ATTACHURLPATH%%/%v" %v/>`,
		image.fileName, image.DimensionHTMLAttributes())

	return []byte(replacedTag)
}

// AllMetaDataHTML returns the Foswiki metadata HTML code for all decoded base64-embedded images.
func (c *imageConverter) AllMetaDataHTML() []byte {
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

// addImage decodes base64 image, saves image to disk and saves the metadata of the image for AllMetaDataHTML().
// It sets a filename with a fixed prefix and a counter.
func (c *imageConverter) addImage(base64Data []byte, fileExtension string) (*base64Image, error) {
	fileName := fmt.Sprintf("base64-image%04d.%v", len(c.images), fileExtension)

	// decode base64 image and extract metadata
	image, err := NewBase64Image(fileName, base64Data)
	if err != nil {
		return nil, err
	}

	// write image to pubdir
	err = c.imageWriter.WriteImage(c.foswikiDoc.PubDir(), image)
	if err != nil {
		return nil, err
	}

	// add image to imageConverter image array
	c.images = append(c.images, *image)

	return image, nil
}
