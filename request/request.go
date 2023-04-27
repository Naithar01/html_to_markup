package request

import (
	"net/http"

	"golang.org/x/net/html"
)

func RequestHttp(url string) (*html.Node, error) {
	// Get HTML
	resp, err := http.Get(url)

	if err != nil {
		resp.Body.Close()

		return nil, err
	}

	defer resp.Body.Close()

	// HTML => ( Parse) => tree HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		resp.Body.Close()

		return nil, err
	}

	return doc, nil
}
