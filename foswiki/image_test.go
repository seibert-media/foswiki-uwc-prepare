package foswiki

import (
	"testing"
)

type imageTest struct {
	foswikiDoc Doc
	imageTags  []imageTag
}

type imageTag struct {
	originHTML   string
	replacedHTML string
}

var imageTests = []imageTest{
	{
		Doc{"/tmp/data/webname1/pagename1.txt"},
		[]imageTag{
			{
				`<img alt="" src="data:image/png;base64,iVBORw0KGgI=" />`,
				`<img alt="" src="%ATTACHURLPATH%/webname1/pagename1/base64-image0000.png" />`,
			},
		},
	},
	{
		Doc{"/tmp/data/WEB_NAME2/PAGE_NAME-2.txt"},
		[]imageTag{
			{`<img width="123" alt="" src="data:image/gif;base64,iVBORw0KGgI=" border="0"/>`,
				`<img alt="" src="%ATTACHURLPATH%/WEB_NAME2/PAGE_NAME-2/base64-image0000.gif" />`,
			},
			{`<img width="123" alt="" src="data:image/jpeg;base64,iVBORw0KGgI=" border="0"/>`,
				`<img alt="" src="%ATTACHURLPATH%/WEB_NAME2/PAGE_NAME-2/base64-image0001.jpeg" />`,
			},
		},
	},
}

func TestReplaceBase64Tag(t *testing.T) {
	for _, imageTest := range imageTests {
		foswikiImage := ImageConverter{Document: imageTest.foswikiDoc}

		for _, imageTag := range imageTest.imageTags {
			originBytes := []byte(imageTag.originHTML)

			if result := foswikiImage.ReplaceBase64Tag(originBytes); string(result) != imageTag.replacedHTML {
				t.Error("For:", imageTag.originHTML, "\nexpected:", imageTag.replacedHTML, "\ngot:", string(result))
			}
		}
	}
}
