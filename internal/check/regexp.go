package check

import "regexp"

var (
	regexImg  = regexp.MustCompile(`!\[(.*?)\]\(([^)]+)\)`)
	regexLink = regexp.MustCompile(`\[(.*?)\]\((https?://[^\s\)]+)\)`)
)
