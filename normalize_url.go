package main

import (
	"fmt"
	"net/url"
	"strings"
)

// changes inputUrl to a consistent format
func normalizeURL(inputUrl string) (string, error) {
	parsedURL, err := url.Parse(inputUrl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}
	parsedURL.Host = strings.ToLower(strings.TrimPrefix(parsedURL.Hostname(), "www."))
	parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/")

	return parsedURL.String(), err
}
