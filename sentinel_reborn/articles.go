package main

import (
	"regexp"
)

func Article(s string) string {

	re := regexp.MustCompile(`\b([Aa])\s+([AEIOUaeiouhH]\w+)`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		parts := re.FindStringSubmatch(match)
		if parts[1] == "A" {
			return "An " + parts[2]
		}
		return "An " + parts[2]

	})

}
