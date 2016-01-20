package foswiki

import "fmt"

// Image provides the metadata of a Foswiki image
type Image struct {
	fileName string
	size     int
	width    int
	height   int
}

// DimensionHTMLAtrtibures returns the width and height attributes of a Foswiki Image as HTML code.
func (i *Image) DimensionHTMLAttributes() string {
	if i.width == 0 || i.height == 0 {
		return ""
	}
	return fmt.Sprintf(`width="%v" height="%v" `, i.width, i.height)
}
