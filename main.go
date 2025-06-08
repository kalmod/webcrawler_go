package main

import (
	"fmt"
	"os"
)

// https://wagslane.dev/

func main() {
	clArgs := os.Args[1:]
	nArgs := len(clArgs)
	if nArgs < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if nArgs > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	var BASE_URL string = clArgs[0]

	fmt.Printf("starting crawl of: \x1b[4;36m%s\x1b[0m\n", BASE_URL)

	visitedPages := make(map[string]int)
	crawlPage(BASE_URL, BASE_URL, visitedPages)
}
