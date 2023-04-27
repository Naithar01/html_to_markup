package main

import (
	"log"
	"net/http"

	"github.com/Naithar01/html_to_markup/fix"
	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}

	body, err := fix.FilterBodyData(doc)
	if err != nil {
		log.Println(err.Error())
	}

	fix.PrintBodyNodeList(body, "")
}
