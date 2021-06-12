package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

var links map[address]*Link

func ParseLinks(r io.Reader) (l map[address]*Link) {
	l = make(map[address]*Link)
	node, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	crawl(node, l)
	return l
}

func crawl(n *html.Node, l map[address]*Link) {
	checkForATypes(n, l)
	if n.NextSibling != nil {
		crawl(n.NextSibling, l)
	}
	if n.FirstChild != nil {
		crawl(n.FirstChild, l	)
	}
}

func checkForATypes(n *html.Node, l map[address]*Link) {
	if n.DataAtom == 1 {
		for _, index := range n.Attr {
			if index.Key == "href" {
				var link Link
				link.Href = index.Val
				link.Text = aCrawler(n.FirstChild)
				l[n] = &link
			}
		}
	}
}

func aCrawler(n *html.Node) (text string) {
	if n.Type == html.TextNode {
		text = n.Data
	}
	return
}

type address *html.Node

type Link struct {
	Href string
	Text string
}
