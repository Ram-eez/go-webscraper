package main

import (
	"fmt"
	"net/http"
	"sync"

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

func getLink(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	if resp, err := http.Get(url); err == nil {
		if rootNode, err := html.Parse(resp.Body); err == nil {
			for _, link := range findLinks(rootNode) {
				fmt.Printf("link : %s\n", link)
			}
		}
	}
}
func main() {

	var wg sync.WaitGroup

	urls := []string{
		"https://en.wikipedia.org/wiki/Web_scraping",
		"https://www.scrapingbee.com/blog/web-scraping-go/",
	}

	// resultsChannel := make(chan string, len(ulrs))

	for _, url := range urls {
		wg.Add(1)
		go getLink(url, &wg)
	}

	wg.Wait()
}
