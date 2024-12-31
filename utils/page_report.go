package utils

import (
	"fmt"
	"sort"
)

type PageReport struct {
	URL   string
	Count int
}

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	var reports []PageReport
	for url, count := range pages {
		reports = append(reports, PageReport{URL: url, Count: count})
	}

	sort.Slice(reports, func(i, j int) bool {
		if reports[i].Count == reports[j].Count {
			return reports[i].URL < reports[j].URL
		}
		return reports[i].Count > reports[j].Count
	})

	for _, report := range reports {
		fmt.Printf("Found %d internal links to %s\n", report.Count, report.URL)
	}
}
