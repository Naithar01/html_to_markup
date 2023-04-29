package main

import (
	"log"

	"github.com/Naithar01/html_to_markup/fix"
	"github.com/Naithar01/html_to_markup/request"
)

func main() {
	doc, err := request.RequestHttp("https://naithar01.tistory.com/168")
	if err != nil {
		log.Println(err.Error())
	}

	body, err := fix.SelectTagElement(doc, "body")

	d, err := fix.SelectClassElement(body, ".wrap-right")

	if d == nil {
		log.Println("zz")
	}

	fix.PrintNodeList(d)

}
