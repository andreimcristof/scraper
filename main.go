package main

import (
	"flag"
	"io/ioutil"
	"log"
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
	urlArgHint := "enter url to parse"
	tagArgHint := "(optional) enter a specific tag to retrieve only"
	urlFlagPtr := flag.String("url", "", urlArgHint)
	tagFlagPtr := flag.String("tag", "", tagArgHint)
	flag.Parse()
	if *urlFlagPtr == "" {
		log.Fatal("must specify --url flag")
	}
	println("parsing url:", *urlFlagPtr, "| tag:", *tagFlagPtr)
	return *urlFlagPtr, *tagFlagPtr
}

func parse(url string, tag string) []TagWithContent {
	response := fetch(url)
	parsed := ParseHTMLPage(response, tag)
	return parsed
}

func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		println("url err", err.Error())
		os.Exit(2)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	parsedResponse := string(body)
	return parsedResponse
}
