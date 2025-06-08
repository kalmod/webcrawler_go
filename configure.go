package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

// safely updated pages map.
// Locks to update and then unlocks
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	isFirst = true
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return !isFirst
	}
	cfg.pages[normalizedURL] = 1
	return isFirst
}

func (cfg *config) overPageLimit() (isOverLimit bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if len(cfg.pages) >= cfg.maxPages {
		return !isOverLimit
	}
	return isOverLimit
}

// Creates and returns new config
func Configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("configure::%s:: coundlt parse url %s", FormattedErrorText(), rawBaseURL)

	}

	return &config{
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, CHANNEL_LIMIT), // buffered channel
		wg:                 &sync.WaitGroup{},
	}, nil
}
