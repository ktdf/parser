package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

func ParseLinks(r io.Reader) (l []Link) {
	node, _ := html.Parse(r)
	fmt.Println(node)
	return nil
}

type Link struct {
	Href string
	Text string
}