package fix

import (
	"fmt"

	"golang.org/x/net/html"
)

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
