package htmlparser

import (
	"io/ioutil"
	"net/http"
	"os"
)

func bodyOfURL(url string) {
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
