package cmd

import "regexp"

var (
	path, imgPath string
	regexImg      = regexp.MustCompile(`!\[(.*?)\]\(([^)]+)\)`)
	regexLink     = regexp.MustCompile(`\[(.*?)\]\((https?://[^\s\)]+)\)`)
)
