package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Retrieve html from provided URL (rawURL)
func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("getHTML::Recieved error code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("getHTML::%s::Content-type = %s", FormattedErrorText(), contentType)
	}

	defer resp.Body.Close()

	htmlBodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(htmlBodyData), nil
}
