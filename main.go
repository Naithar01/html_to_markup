package main

import (
	"fmt"
	"net/http"

	"github.com/Naithar01/html_to_markup/fix"
	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(fix.GetBodyData(doc))
}
