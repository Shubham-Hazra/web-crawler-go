package utils

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	node, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	var urls []string
	extractLinks(node, &urls)

	var resolvedURLs []string
	for _, rawURL := range urls {
		parsedURL, err := url.Parse(rawURL)
		if err != nil {
			continue
		}

		resolvedURLs = append(resolvedURLs, baseURL.ResolveReference(parsedURL).String())
	}

	return resolvedURLs, nil
}

func extractLinks(node *html.Node, urls *[]string) {
	if node.Type == html.ElementNode && node.DataAtom == atom.A {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				*urls = append(*urls, attr.Val)
				break
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		extractLinks(child, urls)
	}
}
