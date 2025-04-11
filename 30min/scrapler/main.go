package main

import (
	"fmt"
	"net/http"
	"path"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

var ULRS_TO_PARSE = []string{"https://go.dev", "https://en.wikipedia.org", "https://github.com/"}

func run() error {
	client := http.Client{Timeout: 2 * time.Second}

	var extractURLs func(n *html.Node, pURL string) (s []string)
	extractURLs = func(n *html.Node, pURL string) (s []string) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && len(attr.Val) > 0 {
					if []rune(attr.Val)[0] == '/' {
						s = append(s, pURL+attr.Val)
					} else {
						pURL = attr.Val
						s = append(s, attr.Val)
					}
				}
			}
		}

		for v := n.FirstChild; v != nil; v = v.NextSibling {
			s = append(s, extractURLs(v, pURL)...)
		}

		return s
	}

	type scrap_func func(c *http.Client, url string, wg *sync.WaitGroup) (err error)
	var scrap scrap_func
	scrap = func(c *http.Client, url string, wg *sync.WaitGroup) (err error) {
		defer wg.Done()

		resp, err := c.Get(url)
		if err != nil {
			return fmt.Errorf("http Get error: %w", err)
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			return fmt.Errorf("html body parse error: %w", err)
		}

		uniquePath := make(map[string]interface{})
		for _, s := range extractURLs(doc, url) {
			uniquePath[s] = struct{}{}
		}

		for v := range uniquePath {
			println(path.Base(url), "->", v)
		}
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(ULRS_TO_PARSE))
	for _, u := range ULRS_TO_PARSE {
		go func(f scrap_func) {
			if err := f(&client, u, wg); err != nil {
				fmt.Printf("URL %s scrap error: %v\n", u, err)
			}
		}(scrap)
	}

	wg.Wait()

	return nil
}
