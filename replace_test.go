package main

import "testing"

var verbatimTests = map[string]string{
	"<verbatim>test</verbatim>":                 "%CODE%test%ENDCODE%",
	"<Verbatim>Test</verbatim>":                 "%CODE%Test%ENDCODE%",
	"<Verbatim>Test</verbatiM>":                 "%CODE%Test%ENDCODE%",
	"<verbatim>Test!@#$%^&*()(*&^%$</verbatim>": "%CODE%Test!@#$%^&*()(*&^%$%ENDCODE%",
	`
	<verbatim>
	Test
	</verbatim>`: `
	%CODE%
	Test
	%ENDCODE%`,
	`
	<verbatim>
	Test1
	</verbatim>
	ABC
	<verbatim>
	Test2
	</verbatim>`: `
	%CODE%
	Test1
	%ENDCODE%
	ABC
	%CODE%
	Test2
	%ENDCODE%`,
	`
	<verbatim>
	Test1
	<verbatim>
	ABC
	</verbatim>`: `
	%CODE%
	Test1
	<verbatim>
	ABC
	%ENDCODE%`,
	`
	<verbatim>
	Test1
	</verbatim>
	ABC
	</verbatim>`: `
	%CODE%
	Test1
	%ENDCODE%
	ABC
	</verbatim>`,
}

func TestReplaceAllVerbatimTags(t *testing.T) {
	for origin, expected := range verbatimTests {

		originBytes := []byte(origin)

		if result := replaceAllVerbatimTags(&originBytes); string(result) != expected {
			t.Error("For:", origin, "\nexpected:", expected, "\ngot:", string(result))
		}
	}
}
