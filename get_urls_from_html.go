package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/url"
	"strings"
)

// obtains slice of urls from html
func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	var urls []string
	traverseParsedHTML(doc, baseURL, &urls)

	return urls, nil
}

func traverseParsedHTML(node *html.Node, baseUrl *url.URL, links *[]string) {

	for n := range node.Descendants() {
		if n.DataAtom == atom.A && n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					href, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("traverseParsedHTML::%s::couldn't parse href '%v': %v\n", FormattedErrorText(), attr.Val, err)
						continue
					}
					resolvedURL := baseUrl.ResolveReference(href)
					*links = append(*links, resolvedURL.String())
				}
			}
		}
	}
}
