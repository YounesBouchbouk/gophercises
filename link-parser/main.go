package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ReadFileAndParseFile(path string) (*html.Node, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	//doc, err := html.
	doc, err := html.Parse(file)

	if err != nil {
		return nil, err
	}

	return doc, nil
}

func main() {

	doc, err := ReadFileAndParseFile("ex2.html")

	if err != nil {
		panic(err)
	}

	links := ProcessFile(doc)
	fmt.Printf("%+v\n", links)

}

func ProcessFile(doc *html.Node) []Link {
	links := []Link{}

	var ProcessLinks func(*html.Node)

	ProcessLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			ProcessNode(doc, &links)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ProcessLinks(c)
		}
	}

	ProcessLinks(doc)
	return links
}

func ProcessNode(n *html.Node, links *[]Link) {
	switch n.Data {
	case "a":
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			for _, a := range n.Attr {
				if a.Key == "href" {
					text := extractText(n)
					*links = append(*links, Link{Href: n.Attr[0].Val, Text: text})

				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ProcessNode(c, links)
	}
}

func extractText(n *html.Node) string {

	var text string
	if n.Type != html.ElementNode && n.Data != "a" && n.Type != html.CommentNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return strings.Trim(text, "\n")
}
