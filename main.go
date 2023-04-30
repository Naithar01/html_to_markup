package main

import (
	"log"

	"github.com/Naithar01/html_to_markup/fix"
	"github.com/Naithar01/html_to_markup/request"
)

func main() {
	doc, err := request.RequestHttp("https://naithar01.tistory.com/")
	if err != nil {
		log.Println(err.Error())
	}

	body, err := fix.SelectTagElement(doc, "body")

	d, err := fix.SelectClassElements(body, ".wrap-right.main.area-main.area-common.article-type-common")

	if d == nil {
		log.Println("zz")
	}

	for _, dd := range d {
		fix.PrintNodeList(dd)

	}

}
