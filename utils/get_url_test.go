package utils

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
		expectErr bool
	}{{
		name:     "absolute and relative URLs with root relative path",
		inputURL: "https://blog.boot.dev",
		inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
		expected:  []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		expectErr: false,
	}, {
		name:     "absolute and relative URLs with trailing slash",
		inputURL: "https://blog.boot.dev/",
		inputBody: `
		<html>
			<body>
				<a href="/">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
		expected:  []string{"https://blog.boot.dev/", "https://other.com/path/one"},
		expectErr: false,
	}, {
		name:      "empty HTML body",
		inputURL:  "https://blog.boot.dev",
		inputBody: ``,
		expected:  []string{},
		expectErr: false,
	}, {
		name:     "invalid base URL",
		inputURL: ":invalid-url",
		inputBody: `
		<html>
			<body>
				<a href="/path/one">Link</a>
			</body>
		</html>
		`,
		expected:  nil,
		expectErr: true,
	}, {
		name:     "invalid href attribute",
		inputURL: "https://blog.boot.dev",
		inputBody: `
		<html>
			<body>
				<a href="::invalid-url">Link</a>
			</body>
		</html>
		`,
		expected:  []string{},
		expectErr: false,
	},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.inputBody, tc.inputURL)
			if (err != nil) != tc.expectErr {
				t.Errorf("Test %v - '%s' FAIL: expected error=%v, got error=%v", i, tc.name, tc.expectErr, err != nil)
				return
			}
			if actual == nil {
				actual = []string{}
			}
			if tc.expected == nil {
				tc.expected = []string{}
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
			}

		})
	}
}
