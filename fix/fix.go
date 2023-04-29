package fix

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func PrintNodeAttr(element_node_attrs []html.Attribute) {
	for _, attr := range element_node_attrs {
		fmt.Printf(` %s="%s"`, attr.Key, attr.Val)
	}
}

// Start Body => ... => nil
func PrintNodeList(element_node *html.Node) {

	var printNodeFunc func(*html.Node, string)

	printNodeFunc = func(element_node *html.Node, indent string) {
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
				printNodeFunc(children_node, indent+"  ")
			}

			fmt.Printf("%s</%s>\n", indent, element_node.Data)

		}

		// 공백 제거 => "  " => len() == 0
		if len(strings.TrimSpace(element_node.Data)) > 0 && element_node.Type == html.TextNode {
			fmt.Printf("%s%s\n", indent, element_node.Data)
		}
	}

	printNodeFunc(element_node, "")
}

func SelectTagElement(html_element *html.Node, tag_name string) (*html.Node, error) {
	var return_element *html.Node

	var select_tag_func func(*html.Node)

	select_tag_func = func(element *html.Node) {
		if element == nil {
			return
		}

		if element.Type == html.ElementNode && element.Data == tag_name && return_element == nil {
			return_element = element
		}

		for children_node := element.FirstChild; children_node != nil; children_node = children_node.NextSibling {
			select_tag_func(children_node)
		}

	}

	select_tag_func(html_element)

	if return_element == nil {
		return nil, errors.New("cant find tag data")
	}

	return return_element, nil
}

func SelectTagElements(html_element *html.Node, tag_name string) ([]*html.Node, error) {
	var return_element []*html.Node

	var select_tag_func func(*html.Node)

	select_tag_func = func(element *html.Node) {
		if element == nil {
			return
		}

		if element.Type == html.ElementNode && element.Data == tag_name {
			return_element = append(return_element, element)
		}

		for children_node := element.FirstChild; children_node != nil; children_node = children_node.NextSibling {
			select_tag_func(children_node)
		}

	}

	select_tag_func(html_element)

	if return_element == nil {
		return nil, errors.New("cant find tag data")
	}

	return return_element, nil
}

func SelectClassElement(html_element *html.Node, selector string) (*html.Node, error) {
	classes := strings.Split(selector, ".")[1:] // 셀렉터에서 클래스 이름만 추출합니다.
	var return_element *html.Node

	var select_element_func func(*html.Node)

	select_element_func = func(element *html.Node) {
		if element == nil {
			return
		}

		if element.Type == html.ElementNode {
			class_names := strings.Fields(element.Attr[0].Val)        // 클래스 이름을 추출합니다.
			if len(class_names) > 0 && class_names[0] == classes[0] { // 클래스 이름이 셀렉터에 지정된 클래스와 일치하는지 확인합니다.
				if len(classes) == 1 { // 셀렉터에서 클래스가 한 개만 지정된 경우에는 바로 반환합니다.
					return_element = element
					return
				} else { // 셀렉터에서 클래스가 여러 개 지정된 경우에는 재귀 호출하여 다음 클래스를 찾습니다.
					next_classes := "." + strings.Join(classes[1:], ".")
					next_element, err := SelectClassElement(element.FirstChild, next_classes)
					if err == nil {
						return_element = next_element
						return
					}
				}
			}
		}

		// 현재 요소에서 일치하는 클래스를 찾지 못한 경우, 자식 요소를 탐색합니다.
		for children_node := element.FirstChild; children_node != nil; children_node = children_node.NextSibling {
			select_element_func(children_node)
		}
	}

	select_element_func(html_element)

	if return_element == nil {
		return nil, errors.New("cant find element with selector")
	}

	return return_element, nil
}
