package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

//	if body, err := io.ReadAll(resp.Body); err == nil {
//		print(string(body))
//	}
func findLinks(n *html.Node) (links []string) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c)...)
	}

	return links
}
func main() {
	if resp, err := http.Get("https://en.wikipedia.org/wiki/Web_scraping"); err == nil {
		if rootNode, err := html.Parse(resp.Body); err == nil {
			for _, link := range findLinks(rootNode) {
				fmt.Printf("link : %s\n", link)
			}
		}
	}
}
