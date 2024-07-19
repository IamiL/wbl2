package main

import "regexp"

func SetPattern(pattern string, ignore bool) *regexp.Regexp {
	var re *regexp.Regexp
	if ignore {
		re = regexp.MustCompile("(?i)" + pattern)
	} else {
		re = regexp.MustCompile(pattern)
	}
	return re
}
