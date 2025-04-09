package main

import (
	"net/http"
	"time"

	"golang.org/x/net/html"
)

const ULR_TO_PARSE = "https://go.dev"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := http.Client{Timeout: 2 * time.Second}

	resp, err := client.Get(ULR_TO_PARSE)
	if err != nil {
		return err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	var extractURLs func(n *html.Node, pURL string) (s []string)
	extractURLs = func(n *html.Node, pURL string) (s []string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					if []rune(attr.Val)[0] == '/' {
						s = append(s, pURL+attr.Val)
					} else {
						pURL = attr.Val
						s = append(s, attr.Val)
					}
					// println(attr.Val)
				}
			}
		}

		for v := n.FirstChild; v != nil; v = v.NextSibling {
			s = append(s, extractURLs(v, pURL)...)
		}

		return s
	}

	for _, v := range extractURLs(doc, ULR_TO_PARSE) {
		println(v)
	}

	return nil
}
