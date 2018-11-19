package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	tagsWithContent := parse("https://www.google.se")

	for i := 0; i < len(tagsWithContent); i++ {
		println("content", tagsWithContent[i].content)
		print("tag", tagsWithContent[i].tag)
	}

}

func parse(url string) []TagWithContent {
	response := fetch(url)
	parsed := ParseHTML(response)
	return parsed
}

func fetch(url string) string {
	res, getErr := http.Get(url)
	if getErr != nil {
		os.Exit(1)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	parsedResponse := string(body)
	return parsedResponse
}
