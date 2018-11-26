package scraper

import (
	"strings"
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

func TestParseHTMLPage(t *testing.T) {
	mock := `<html>
				<body>
					<div></div>
				</body>
			</html>`

	expectedTag := "html"
	result := parseHTMLPage(mock, expectedTag)
	tag := result[0]
	hasEmptySpaces := strings.Contains(tag.Content, " ")
	if hasEmptySpaces {
		t.Errorf("parsed content should not contain empty spaces")
	}

	isCorrectTagExracted := strings.EqualFold(tag.Tag, expectedTag)
	if !isCorrectTagExracted {
		t.Errorf("incorrect tag extracted, asked for %v and got instead %v", expectedTag, tag.Tag)
	}

	containsChildTag := strings.Contains(tag.Content, expectedTag)
	if !containsChildTag {
		t.Errorf("parsed content should contain child tags")
	}

	t.Log(tag.Content)
}
