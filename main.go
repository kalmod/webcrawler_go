package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// https://wagslane.dev/

var (
	CHANNEL_LIMIT int = 3
	MAX_PAGES     int = 5
)

func main() {
	clArgs := os.Args[1:]
	nArgs := len(clArgs)
	if nArgs < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	var err error
	if nArgs > 1 {
		CHANNEL_LIMIT, err = strconv.Atoi(clArgs[1])
		if err != nil {
			fmt.Printf("main::%s::could not convert to int > %s", FormattedErrorText(), err.Error())
		}
	}
	if nArgs > 2 {
		MAX_PAGES, err = strconv.Atoi(clArgs[2])
		if err != nil {
			fmt.Printf("main::%s::could not convert to int > %s", FormattedErrorText(), err.Error())
		}
	}

	var BASE_URL string = clArgs[0]

	fmt.Printf("starting crawl of: \x1b[4;36m%s\x1b[0m\n", BASE_URL)

	cfg, err := Configure(BASE_URL, CHANNEL_LIMIT, MAX_PAGES)
	if err != nil {
		os.Exit(1)
	}

	start := time.Now()

	cfg.wg.Add(1)
	go cfg.crawlPage(BASE_URL)
	cfg.wg.Wait() // add and remove from wait group.

	t := time.Now()
	elapsedTime := t.Sub(start)
	fmt.Println("DONE CRAWLING (in my skin)")
	fmt.Printf("\x1b[1;35mELAPSED TIME: %.2f SECONDS\x1b[0m\n", elapsedTime.Seconds())
}
