package main

import (
	"fmt"
	"net/http"
	"strings"
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

				if hasValidExtension(attr.Val) {
					links = append(links, attr.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, findLinks(c)...)
	}

	return links
}

func hasValidExtension(link string) bool {
	validExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".docx", ".mp3", ".mp4"}

	for _, extensions := range validExtensions {
		if strings.HasSuffix(link, extensions) {
			return true
		}
	}
	return false
}

func getLink(url string, wg *sync.WaitGroup, semaphore chan struct{}) {

	defer wg.Done()
	semaphore <- struct{}{}
	fmt.Printf("Starting %s\n", url)
	if resp, err := http.Get(url); err == nil {
		if rootNode, err := html.Parse(resp.Body); err == nil {
			for _, link := range findLinks(rootNode) {
				fmt.Printf("link : %s\n", link)
			}
		}
	}
	<-semaphore
	fmt.Printf("Finished %s\n", url)
}
func main() {

	var wg sync.WaitGroup

	concurrencyLimit := 2
	// here using struct for the semaphore we eleminate the need to occupy any memory
	// here semaphore is a buffered channel
	semaphore := make(chan struct{}, concurrencyLimit)

	urls := []string{
		"https://en.wikipedia.org/wiki/Web_scraping",
		"https://www.scrapingbee.com/blog/web-scraping-go/",
		"https://en.wikipedia.org/wiki/Wiki",
	}

	// resultsChannel := make(chan string, len(ulrs))

	for _, url := range urls {
		wg.Add(1)
		go getLink(url, &wg, semaphore)
	}

	wg.Wait()
}
