package main

import (
	"fmt"
	"sort"
)

type pageInfo struct {
	urlCount int
	URL      string
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`=============================
  REPORT for %s
=============================
`, baseURL)

	var allPages []pageInfo
	for key, val := range pages {
		allPages = append(allPages, pageInfo{val, key})
	}

	sort.Slice(allPages, func(i, j int) bool {
		return allPages[i].urlCount > allPages[j].urlCount
	})

	for _, p := range allPages {
		fmt.Printf("Found \x1b[33m%d\x1b[0m internal links to \x1b[4;34m%s\x1b[0m\n", p.urlCount, p.URL)
	}

}
