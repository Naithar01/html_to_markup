package fix

import (
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

func FilterBodyData(doc *html.Node) (*html.Node, error) {
	var body *html.Node

	var findBodyFunc func(*html.Node)

	findBodyFunc = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}

		for children_node := node.FirstChild; children_node != nil; children_node = children_node.NextSibling {
			findBodyFunc(children_node)
		}
	}

	findBodyFunc(doc)

	if body == nil {
		return nil, errors.New("no find body")
	}

	return body, nil
}

func GetBodyData(doc *html.Node) string {
	var bodyText string
	bodyTagStarted := false

	var getNodeText func(*html.Node)
	getNodeText = func(n *html.Node) {
		if n.Data == "body" && n.Type == 3 {
			bodyTagStarted = !bodyTagStarted
		}

		if n.Type == 3 && bodyTagStarted {
			bodyText += fmt.Sprintf("<%s", n.Data)
			for _, attr := range n.Attr { // class, id 등
				bodyText += fmt.Sprintf(" %s=\"%s\"", attr.Key, attr.Val)
			}
			bodyText += ">"

		} else if n.Type == html.TextNode && bodyTagStarted {
			bodyText += n.Data
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getNodeText(c)
		}

		if n.Type == 3 && bodyTagStarted && n.Data != "html" {
			bodyText += fmt.Sprintf("</%s>", n.Data)
		}
	}

	getNodeText(doc)

	return bodyText
}
