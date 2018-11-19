package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	parse("https://www.google.se")
	// print(r)
}

func parse(url string) string {
	response := fetch(url)
	parsed := ParseHTML(response)
	return parsed.Data
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
