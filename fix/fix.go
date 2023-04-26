package fix

import (
	"fmt"

	"golang.org/x/net/html"
)

func FilterBodyData(doc *html.Node) (*html.Node, error) {
	var body *html.Node

	// doc 변수에 있는 태그를 순회 ( 트리구조 )
	// body 태그가 나오면 변수에 넣고 반환
	// 재귀 함수 형식이기에 함수 변수를 선언
	var findBodyFunc func(*html.Node)

	findBodyFunc = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node

			return
		}

		for node_content := node.FirstChild; node_content != nil; node_content = node.LastChild {
			findBodyFunc(node)
		}
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
