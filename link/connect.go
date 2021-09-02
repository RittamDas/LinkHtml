package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Links struct {
	Href string
	Text string
}

func Connect(r io.Reader) ([]Links, error) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("k")
		return nil, err
	}
	return parseLinks(doc), nil
}

func parseLinks(m *html.Node) []Links {
	var l []Links
	if m.Type == html.ElementNode && m.Data == "a" {
		for _, a := range m.Attr {
			if a.Key == "href" {
				t := parseText(m)
				l = append(l, Links{a.Val, t})
			}
		}
	}
	for k := m.FirstChild; k != nil; k = k.NextSibling {
		allLinks := parseLinks(k)
		l = append(l, allLinks...)
	}
	return l
}

func parseText(m *html.Node) string {
	var t string
	if m.Type != html.ElementNode && m.Data != "a" && m.Type != html.CommentNode {
		t = m.Data
	}
	for k := m.FirstChild; m != nil; m = m.NextSibling {
		t += parseText(k)
	}
	return strings.Trim(t, "\n")
}
