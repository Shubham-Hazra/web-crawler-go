package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/Shubham-Hazra/web-crawler-go/utils"
)

type config struct {
	maxPages           int
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[1]
	maxConcurrency, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatalf("Invalid argument maxConcurrency: %v\n", err)
	}
	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatalf("Invalid argument maxPages: %v\n", err)
	}

	fmt.Println("starting crawl of: " + baseURL)

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Invalid base URL: %v\n", err)
	}

	cfg := &config{
		maxPages:           maxPages,
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go func() {
		cfg.concurrencyControl <- struct{}{}
		cfg.crawlPage(baseURL)
	}()

	cfg.wg.Wait()

	utils.PrintReport(cfg.pages, baseURL)
}
