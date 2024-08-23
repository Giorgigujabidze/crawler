package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseUrl            *url.URL
	mu                 *sync.Mutex
	maxConcurrency     int
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
	pagesCrawled       int
}

func main() {
	cfg := &config{
		pages:              map[string]int{},
		baseUrl:            &url.URL{},
		mu:                 &sync.Mutex{},
		maxConcurrency:     0,
		concurrencyControl: nil,
		wg:                 &sync.WaitGroup{},
		maxPages:           0,
		pagesCrawled:       0,
	}

	args := os.Args[1:]
	cfg.maxConcurrency, _ = strconv.Atoi(args[1])
	cfg.maxPages, _ = strconv.Atoi(args[2])
	cfg.concurrencyControl = make(chan struct{}, cfg.maxConcurrency)
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Println("starting crawl")
	_, err := getHtml(args[0])
	if err != nil {
		fmt.Println("Error getting html")
		os.Exit(1)
	}
	cfg.baseUrl, _ = url.Parse(args[0])

	cfg.crawlPage(args[0])

	cfg.wg.Wait()
	printReport(cfg.pages, cfg.baseUrl.String())
}
