package utils

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme https",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		}, {
			name:     "remove scheme http",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		}, {
			name:     "remove last / https",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		}, {
			name:     "remove last / http",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove last / http",
			inputURL: "http://blog.boot.dev/path/v1/assets/",
			expected: "blog.boot.dev/path/v1/assets",
		},
		{
			name:     "remove last / http",
			inputURL: "http://blog.boot.dev/path/v1/assets",
			expected: "blog.boot.dev/path/v1/assets",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
