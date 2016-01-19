package foswiki

import (
	"os"
	"path/filepath"
)

type Dir struct {
	HomePath string
}

var FOSWIKI_SUBDIRS = map[string]string{
	"DATA": "data",
	"PUB":  "pub",
}

// SubdirsExists check if the need Foswiki subdirectories exists.
func (d *Dir) SubdirsExists() (bool, error) {
	for _, subdir := range FOSWIKI_SUBDIRS {
		subdir = filepath.Clean(d.HomePath + "/" + subdir)
		if _, err := os.Stat(subdir); err != nil {
			return false, err
		}
	}
	return true, nil
}

// DataPath returns the file path to the Foswiki data directory.
func (d *Dir) DataPath() string {
	return filepath.Clean(d.HomePath + "/" + FOSWIKI_SUBDIRS["DATA"])
}

// PubPath returns the file path to the Foswiki pub directory.
func (d *Dir) PubPath() string {
	return filepath.Clean(d.HomePath + "/" + FOSWIKI_SUBDIRS["PUB"])
}
