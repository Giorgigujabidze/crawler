package main

import (
	"fmt"
	"strings"
)

func (cfg *config) addPageVisit(normalizedUrl string) (isFirst bool) {

	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, ok := cfg.pages[normalizedUrl]
	if !ok {
		cfg.pages[normalizedUrl] = 1
		return true
	}
	cfg.pages[normalizedUrl]++
	return false
}

func (cfg *config) checkMaxPagesExceeded() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if cfg.pagesCrawled >= cfg.maxPages {
		return true
	}
	return false
}

func (cfg *config) crawlPage(rawCurrentUrl string) {
	cfg.concurrencyControl <- struct{}{}
	cfg.wg.Add(1)
	defer func() {
		cfg.wg.Done()
		<-cfg.concurrencyControl
	}()

	if cfg.checkMaxPagesExceeded() {
		return
	}

	if !strings.Contains(rawCurrentUrl, cfg.baseUrl.String()) {
		return
	}

	normalRawCurrentUrl, err := normalizeUrl(rawCurrentUrl)
	if err != nil {
		fmt.Println("Error normalizing URL:", err)
		return
	}

	isFirst := cfg.addPageVisit(normalRawCurrentUrl)
	if !isFirst {
		return
	}
	cfg.mu.Lock()
	cfg.pagesCrawled++
	cfg.mu.Unlock()

	fmt.Println("Crawling page", rawCurrentUrl)

	data, err := getHtml(rawCurrentUrl)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}
	urls, err := cfg.getUrlsFromHtml(data, cfg.baseUrl.String())
	if err != nil {
		fmt.Println("Error extracting URLs:", err)
		return
	}

	for _, url := range urls {

		go cfg.crawlPage(url)

	}
}
