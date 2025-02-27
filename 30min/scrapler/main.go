package main

import (
	"net/http"

	"golang.org/x/net/html"
)

const ULR_TO_PARSE = "https://go.dev"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	resp, err := http.Get(ULR_TO_PARSE)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	s, err := extractURL(doc)
	if err != nil {
		return err
	}

	for _, v := range s {
		println(v)
	}

	return nil
}

func extractURL(n *html.Node) (s []string, err error) {
	if n.Data == "a" {
		println(n.Attr)
	}
	// for {

	// }

	return
}
