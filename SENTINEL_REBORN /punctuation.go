package main

import (
	"regexp"
)

func punctuation(punc string) string {

	re1 := regexp.MustCompile(`\s+([.,!?:;]+)`)
	punc = re1.ReplaceAllString(punc, "$1")

	re2 := regexp.MustCompile(`([.,!?:;]+)`)
	punc = re2.ReplaceAllString(punc, "$1 $2")

	return punc
}
