package crawl

import (
	"net/http"
	"path"
	"os"
	"io"
	"fmt"
	"log"
	"golang.org/x/net/html"
)

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}


func findLinks(url string) ([] string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return visit(nil, doc), nil
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	filename = path.Base(resp.Request.URL.Path)
	log.Printf(filename)
	if filename == "/" {
		filename = "index.html"
	}
	f, err := os.Create("static/crawl_" + filename)
	if err != nil {
		return
	}
	n, err = io.Copy(f, resp.Body)
	if err == nil {
		err = f.Close()
	}
	return
}

func Crawl(url string) {
	log.Printf(url)
	if links, err := findLinks(url); err == nil {
		log.Printf(links[0])
		for _, link := range links {
			log.Printf(link)
			fetch(link)
		}
	} else {
		fmt.Printf("%v", err)
	}
}