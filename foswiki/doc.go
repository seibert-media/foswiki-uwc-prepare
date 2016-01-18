package foswiki

import (
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
)

type Doc struct {
	Path string
}

func (d Doc) PageName() string {
	// pagename = filename - extension
	return strings.Replace(filepath.Base(d.Path), filepath.Ext(d.Path), "", 1)
}

func (d Doc) WebName() string {
	var pathReg = regexp.MustCompile(`.*/(.+?)/.+?\.txt`)
	unixPath := filepath.ToSlash(d.Path)
	matchedParts := pathReg.FindStringSubmatch(unixPath)

	if matchedParts == nil {
		return ""
	}
	return matchedParts[1]
}

func (d Doc) PubDir() string {
	var pathReg = regexp.MustCompile(`(.*)/data/(.+?)/(.+?)\.txt`)
	unixPath := filepath.ToSlash(d.Path)
	matchedParts := pathReg.FindStringSubmatch(unixPath)

	if matchedParts == nil {
		return ""
	}
	pubDir := fmt.Sprintf("%v/pub/%v/%v/", matchedParts[1], matchedParts[2], matchedParts[3])
	return filepath.Clean(pubDir)
}
