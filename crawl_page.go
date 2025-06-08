package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.overPageLimit() {
		return
	}

	fmt.Printf("\t....crawling \x1b[4;34m%s\x1b[0m\n", rawCurrentURL)
	currentURL, currentURL_err := url.Parse(rawCurrentURL)
	if currentURL_err != nil {
		fmt.Printf("\tcrawlPage::%v::Error trying to parse urls", FormattedErrorText())
		fmt.Printf("\t\tBase URL %v > %v", FormattedErrorText(), currentURL_err.Error())
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		fmt.Printf("\t\tcrawlPage::%v::Domains don't match > %s != %s\n", FormattedErrorText(), cfg.baseURL.Hostname(), currentURL.Hostname())
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("\t\tcrawlPage::%v::Error normalizing url > %s\n", FormattedErrorText(), err.Error())
		return
	}

	// sagfe update
	if !cfg.addPageVisit(normalizedCurrentURL) {
		fmt.Printf("\t\t ● %s - \x1b[33mVisited\x1b[0m\n", normalizedCurrentURL)
		return
	}

	htmlBody, err := getHTML(normalizedCurrentURL)
	if err != nil {
		fmt.Printf("\t\tcrawlPage::%v::Could not get HTML > %s\n", FormattedErrorText(), err.Error())
		return
	}

	urls, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("\t\tcrawlPage::%v::Could not build slice of url links > %s\n", FormattedErrorText(), err.Error())
		return
	}

	for _, nextURL := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
	return
}

// func basicCrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
// 	fmt.Printf("\t....crawling \x1b[4;34m%s\x1b[0m\n", rawCurrentURL)
// 	baseURL, baseURL_err := url.Parse(rawBaseURL)
// 	currentURL, currentURL_err := url.Parse(rawCurrentURL)
// 	if baseURL_err != nil || currentURL_err != nil {
// 		fmt.Printf("\tcrawlPage::%v::Error trying to parse urls", FormattedErrorText())
// 		fmt.Printf("\t\tBase URL %v > %v", FormattedErrorText(), baseURL_err.Error())
// 		fmt.Printf("\t\tBase URL %v > %v", FormattedErrorText(), currentURL_err.Error())
// 		return
// 	}
//
// 	if baseURL.Hostname() != currentURL.Hostname() {
// 		fmt.Printf("\tcrawlPage::%v::Domains don't match > %s != %s\n", FormattedErrorText(), baseURL.Hostname(), currentURL.Hostname())
// 		return
// 	}
//
// 	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
// 	if err != nil {
// 		fmt.Printf("\tcrawlPage::%v::Error normalizing url > %s\n", FormattedErrorText(), err.Error())
// 		return
// 	}
//
// 	if _, exists := pages[normalizedCurrentURL]; exists {
// 		pages[normalizedCurrentURL]++
// 		fmt.Printf("\t\t ● %s - \x1b[33mVisited\x1b[0m\n", normalizedCurrentURL)
// 		return
// 	} else {
// 		pages[normalizedCurrentURL] = 1
// 	}
//
// 	htmlBody, err := getHTML(normalizedCurrentURL)
// 	if err != nil {
// 		fmt.Printf("\tcrawlPage::%v::Could not get HTML > %s\n", FormattedErrorText(), err.Error())
// 		return
// 	}
//
// 	urls, err := getURLsFromHTML(htmlBody, normalizedCurrentURL)
// 	if err != nil {
// 		fmt.Printf("\tcrawlPage::%v::Could not build slice of url links > %s\n", FormattedErrorText(), err.Error())
// 		return
// 	}
//
// 	for _, nextURL := range urls {
// 		basicCrawlPage(rawBaseURL, nextURL, pages)
// 	}
// 	return
// }
