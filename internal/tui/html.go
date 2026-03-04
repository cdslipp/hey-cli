package tui

import "hey-cli/internal/htmlutil"

func htmlToText(s string) string {
	return htmlutil.ToText(s)
}

func extractImageURLs(s string) []string {
	return htmlutil.ExtractImageURLs(s)
}
