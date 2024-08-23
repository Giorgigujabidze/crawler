package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func (cfg *config) getUrlsFromHtml(htmlBody, rawBaseUrl string) ([]string, error) {

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		fmt.Println("Error parsing HTML")
		return nil, err
	}
	var urls []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					appendVal := a.Val
					if !strings.Contains(a.Val, "http") {
						appendVal, _ = url.JoinPath(rawBaseUrl, a.Val)
					}
					urls = append(urls, appendVal)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls, nil
}
