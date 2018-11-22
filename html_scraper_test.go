package scraper

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tc := TagWithContent{
		Content: "",
		Tag:     "",
	}
	actual := tc.isValidTag("")
	if actual {
		t.Errorf("shold fail when no tag and no content")
	}

	tc = TagWithContent{
		Content: "content",
		Tag:     "",
	}
	actual = tc.isValidTag("")
	if actual {
		t.Errorf("shold fail when no tag")
	}

	tc = TagWithContent{
		Content: "",
		Tag:     "div",
	}
	actual = tc.isValidTag("")
	if actual {
		t.Errorf("shold fail when no content")
	}

	tc = TagWithContent{
		Content: "content",
		Tag:     "div",
	}
	actual = tc.isValidTag("")
	if !actual {
		t.Errorf("shold be ok when all passed and has no specifictag")
	}

	tc = TagWithContent{
		Content: "",
		Tag:     "div",
	}
	actual = tc.isValidTag("div")
	if actual {
		t.Errorf("shold fail when no content")
	}

	tc = TagWithContent{
		Content: "content",
		Tag:     "div",
	}
	actual = tc.isValidTag("div")
	if !actual {
		t.Errorf("shold be ok when all passed and has specifictag")
	}
}
