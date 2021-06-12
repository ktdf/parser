package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

var links []*Link

func ParseLinks(r io.Reader) []*Link {
	l := &links
	node, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	crawl(node, l)
	return links
}

func crawl(n *html.Node, l *[]*Link) {
	checkForATypes(n, l)
	if n.NextSibling != nil {
		crawl(n.NextSibling, l)
	}
	if n.FirstChild != nil {
		crawl(n.FirstChild, l)
	}
}

func checkForATypes(n *html.Node, l *[]*Link) {
	if n.DataAtom == 1 {
		for _, index := range n.Attr {
			if index.Key == "href" {
				var link Link
				link.Href = index.Val
				link.Text = *aCrawler(n.FirstChild, &link.Text)
				*l = append(*l, &link)
			}
		}
	}
}

func aCrawler(n *html.Node, text *string)  *string {
	var tmpText string
	if n.Type == html.TextNode {
		*text = *text + " " + n.Data
	}
	if n.NextSibling != nil {
		tmpText = *aCrawler(n.NextSibling, &tmpText)
		if tmpText != "" {
			*text = *text + " " + tmpText
		}
	}
	if n.FirstChild != nil {
		tmpText = *aCrawler(n.FirstChild, &tmpText)
		if tmpText != "" {
			 *text = *text + " " + tmpText
		}
	}
	return text
}

type address *html.Node

type Link struct {
	Href string
	Text string
}
