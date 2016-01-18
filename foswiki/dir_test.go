package foswiki

import (
	"testing"
	"path/filepath"
)

func TestDataPath(t *testing.T) {
	path := "/path/to/foswiki/"
	expected := filepath.Clean(path + "/data")
	d := Dir{path}
	if d.DataPath() != expected {
		t.Error("For:", path, "expected:", expected, "got:", d.HomePath)
	}
}

func TestPubPath(t *testing.T) {
	path := "/path/to/foswiki/"
	expected := filepath.Clean(path + "/pub")
	d := Dir{path}
	if d.PubPath() != expected {
		t.Error("For:", path, "expected:", expected, "got:", d.HomePath)
	}
}
