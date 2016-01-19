package foswiki

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

type Doc struct {
	Path string
}

// PageName returns the page name of the current Foswiki document.
func (d Doc) PageName() string {
	// pagename = filename - extension
	return strings.Replace(filepath.Base(d.Path), filepath.Ext(d.Path), "", 1)
}

// WebName returns the web name of the current Foswiki document.
func (d Doc) WebName() string {
	var pathReg = regexp.MustCompile(`.*/(.+?)/.+?\.txt`)
	unixPath := filepath.ToSlash(d.Path)
	matchedParts := pathReg.FindStringSubmatch(unixPath)

	if matchedParts == nil {
		return ""
	}
	return matchedParts[1]
}

// PubDir returns the path to the Foswiki pub folder for the current Foswiki document.
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
