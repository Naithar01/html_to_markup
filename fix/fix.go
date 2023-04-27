package fix

import (
	"errors"
	"fmt"
	"strings"

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

// Start Body => ... => nil
func PrintBodyNodeList(element_node *html.Node, indent string) {
	if element_node == nil {
		return
	}

	// element가 Text인 경우는
	// 또 하나의 태그일 경우
	if element_node.Type == html.ElementNode {
		fmt.Printf("%s<%s", indent, element_node.Data)

		PrintNodeAttr(element_node.Attr)

		fmt.Println(">")

		for children_node := element_node.FirstChild; children_node != nil; children_node = children_node.NextSibling {
			// 단계별로 재귀를 부르기 때문에 공백의 크기가 점점 커짐
			PrintBodyNodeList(children_node, indent+"  ")
		}

		fmt.Printf("%s</%s>\n", indent, element_node.Data)

	}

	// 공백 제거 => "  " => len() == 0
	if len(strings.TrimSpace(element_node.Data)) > 0 && element_node.Type == html.TextNode {
		fmt.Printf("%s%s\n", indent, element_node.Data)
	}
}

func PrintNodeAttr(element_node_attrs []html.Attribute) {
	for _, attr := range element_node_attrs {
		fmt.Printf(` %s="%s"`, attr.Key, attr.Val)
	}
}
