package main

import (
	"io"
	"net/http"
)

func main() {
	if resp, err := http.Get("https://en.wikipedia.org/wiki/Web_scraping"); err == nil {
		defer resp.Body.Close()
		if body, err := io.ReadAll(resp.Body); err == nil {
			print(string(body))
		}
	}
}
