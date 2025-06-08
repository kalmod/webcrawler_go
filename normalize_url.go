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
	var fullPath string
	fullPath = parsedURL.Host + parsedURL.Path
	fullPath = strings.TrimPrefix(fullPath, "www.")
	fullPath = strings.TrimSuffix(fullPath, "/")

	normalizedURL := strings.ToLower(fullPath)

	return normalizedURL, err
}
