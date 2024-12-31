package utils

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func GetHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode >= 400 {
		return "", errors.New("recieved a status code of more than 400")
	} else if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return "", errors.New("content type is not text/html")
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}
