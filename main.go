package main

import (
	"log"

	"github.com/Naithar01/html_to_markup/request"
)

func main() {
	doc, err := request.RequestHttp("https://example.com/")
	if err != nil {
		log.Println(err.Error())
	}

}
