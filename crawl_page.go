package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Shubham-Hazra/web-crawler-go/utils"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; !exists {
		cfg.pages[normalizedURL] = 1
		return true
	}
	cfg.pages[normalizedURL]++
	return false
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	normalizedCurrentURL, err := utils.NormalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("Error normalizing URL %s: %v\n", rawCurrentURL, err)
		return
	}

	if !cfg.addPageVisit(normalizedCurrentURL) {
		return
	}

	fmt.Printf("Crawling URL: %s\n", rawCurrentURL)
	html, err := utils.GetHTML(rawCurrentURL)
	if err != nil {
		log.Printf("Error fetching HTML for URL %s: %v\n", rawCurrentURL, err)
		return
	}

	urls, err := utils.GetURLsFromHTML(html, rawCurrentURL)
	if err != nil {
		log.Printf("Error extracting URLs from HTML for URL %s: %v\n", rawCurrentURL, err)
		return
	}

	for _, extractedURL := range urls {
		parsedBaseURL := cfg.baseURL
		parsedCurrentURL, err := url.Parse(extractedURL)
		if err != nil {
			log.Printf("Error parsing extracted URL %s: %v\n", extractedURL, err)
			continue
		}

		if parsedBaseURL.Host != parsedCurrentURL.Host {
			continue
		}

		cfg.wg.Add(1)
		go func(url string) {
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(url)
		}(extractedURL)
	}
}
