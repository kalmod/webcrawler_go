package main

import (
	"fmt"
	"os"
)

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

	html, err := getHTML(BASE_URL)
	if err != nil {
		fmt.Printf("MAIN::%s::using getHTML: %v\n", FormattedErrorText(), err.Error())
		os.Exit(1)
	}

	fmt.Println(html)
}
