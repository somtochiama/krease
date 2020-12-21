package issue

import (
	"io/ioutil"
	"testing"
)

func TestTemplate(t *testing.T) {
	expectedCommentBytes, err := ioutil.ReadFile("../testdata/expected.txt")
	if err != nil {
		t.Errorf("Unexpected error when parsing exepected file: %v", err)
	}
	expectedComment := string(expectedCommentBytes)
	tpl, err := CreateTemplate("../testdata/comment.txt")
	if err != nil {
		t.Errorf("Unexpected error when parsing file: %v", err)
	}

	comment, err := ParseTemplate(tpl, "Somtochi")
	if err != nil {
		t.Errorf("Unexpected error when parsing template: %v", err)
	}

	if expectedComment != comment {
		t.Errorf("Expecting %v, got %v", expectedComment, comment)
	}
}
