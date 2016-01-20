package main

import (
	"regexp"

	"github.com/seibert-media/foswiki-uwc-prepare/foswiki"
)

// replaceAllVerbatimTags decodes base64 embedded images and rewrite the HTML code.
func replaceAllBase64Images(fileContent *[]byte, path string) []byte {
	foswikiDoc := foswiki.Doc{path}
	imageWriter := foswiki.NewBase64ImageWriter()
	foswikiImageConverter := foswiki.NewImageConverter(foswikiDoc, imageWriter)

	var imageTagRe = regexp.MustCompile(`(?si)<img .*?src="data:image/.+?;base64,.+?".*?/>`)
	replacedHTML := imageTagRe.ReplaceAllFunc(*fileContent, foswikiImageConverter.ReplaceBase64Tag)

	return append(replacedHTML, foswikiImageConverter.AllMetaDataHTML()...)
}

// replaceAllVerbatimTags replaces all <verbatim>..</verbatim> tags to %CODE%..%ENDCODE%.
func replaceAllVerbatimTags(fileContent *[]byte) []byte {
	var verbatimTagRe = regexp.MustCompile(`(?si)<verbatim[^\>]*?>(.+?)</verbatim>`)
	codeTag := []byte("%CODE%$1%ENDCODE%")
	return verbatimTagRe.ReplaceAll(*fileContent, codeTag)
}
