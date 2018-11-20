package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url, tag := setupFlags()

	tagsWithContent := parse(url, tag)

	for i := 0; i < len(tagsWithContent); i++ {
		println(tagsWithContent[i].tag, ":", tagsWithContent[i].content)
	}
}

func setupFlags() (string, string) {
	urlFlagPtr := flag.String("url", "https://google.de", "enter url to parse")
	tagFlagPtr := flag.String("tag", "div", "(optional) enter a specific tag to retrieve only")
	flag.Parse()
	println("parsing url:", *urlFlagPtr, "| tag:", *tagFlagPtr)
	return *urlFlagPtr, *tagFlagPtr
}

func parse(url string, tag string) []TagWithContent {
	response := fetch(url)
	parsed := ParseHTMLPage(response, tag)
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
