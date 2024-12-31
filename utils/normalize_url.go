package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	normalizedPath := strings.TrimSuffix(parsedURL.Path, "/")

	cleanedURL := fmt.Sprintf("%s%s", parsedURL.Host, normalizedPath)
	return cleanedURL, nil
}
