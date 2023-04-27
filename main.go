package main

import (
	"log"

	"github.com/Naithar01/html_to_markup/fix"
	"github.com/Naithar01/html_to_markup/request"
)

func main() {

	doc, err := request.RequestHttp("https://example.com/")

	body, err := fix.FilterBodyData(doc)
	if err != nil {
		log.Println(err.Error())
	}

	fix.PrintBodyNodeList(body, "")
}
