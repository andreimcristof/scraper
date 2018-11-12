package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	strRes := fetch("https://google.se")
	parse(strRes)
}

func fetch(url string) string {
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	parsedResponse := string(body)
	return parsedResponse
}

func parse(page string) {
	reader := strings.NewReader(page)
	readerToken := html.NewTokenizer(reader)

	for {
		n := readerToken.Next()
		if n == html.ErrorToken {
			fmt.Println("error")
			return
		}
		fmt.Print("current token: ", n.String())
	}
}
