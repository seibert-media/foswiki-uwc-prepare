package foswiki

import "testing"

func TestPageName(t *testing.T) {
	path := "/path/to/foswiki/data/WEB_NAME/PAGE_NAME-1.txt"
	expected := "PAGE_NAME-1"
	tested := Doc{path}.PageName()
	if tested != expected {
		t.Error("For:", path, "expected:", expected, "got:", tested)
	}
}

func TestWebName(t *testing.T) {
	path := "/path/to/foswiki/data/WEB_NAME/PAGE_NAME-1.txt"
	expected := "WEB_NAME"
	tested := Doc{path}.WebName()
	if tested != expected {
		t.Error("For:", path, "expected:", expected, "got:", tested)
	}
}

func TestPubDir(t *testing.T) {
	path := "/path/to/foswiki/data/WEB_NAME/PAGE_NAME-1.txt"
	expected := "/path/to/foswiki/pub/WEB_NAME/PAGE_NAME-1"
	tested := Doc{path}.PubDir()
	if tested != expected {
		t.Error("For:", path, "expected:", expected, "got:", tested)
	}
}