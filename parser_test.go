package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tc := TagWithContent{
		content: "",
		tag:     "",
	}
	actual := tc.isValidTag("")
	if actual {
		t.Errorf("shold fail when no tag and no content")
	}

	tc = TagWithContent{
		content: "content",
		tag:     "",
	}
	actual = tc.isValidTag("")
	if actual {
		t.Errorf("shold fail when no tag")
	}

	tc = TagWithContent{
		content: "",
		tag:     "div",
	}
	actual = tc.isValidTag("")
	if actual {
		t.Errorf("shold fail when no content")
	}

	tc = TagWithContent{
		content: "content",
		tag:     "div",
	}
	actual = tc.isValidTag("")
	if !actual {
		t.Errorf("shold be ok when all passed and has no specifictag")
	}

	tc = TagWithContent{
		content: "",
		tag:     "div",
	}
	actual = tc.isValidTag("div")
	if actual {
		t.Errorf("shold fail when no content")
	}

	tc = TagWithContent{
		content: "content",
		tag:     "div",
	}
	actual = tc.isValidTag("div")
	if !actual {
		t.Errorf("shold be ok when all passed and has specifictag")
	}
}
