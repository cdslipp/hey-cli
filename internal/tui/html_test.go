package tui

import (
	"testing"
)

// htmlToText and extractImageURLs are thin wrappers around htmlutil.
// Full tests are in internal/htmlutil/htmlutil_test.go.

func TestHtmlToTextDelegates(t *testing.T) {
	got := htmlToText("<p>hello</p>")
	if got != "hello" {
		t.Errorf("htmlToText = %q, want %q", got, "hello")
	}
}

func TestExtractImageURLsDelegates(t *testing.T) {
	urls := extractImageURLs(`<img src="test.png">`)
	if len(urls) != 1 || urls[0] != "test.png" {
		t.Errorf("extractImageURLs = %v, want [test.png]", urls)
	}
}
