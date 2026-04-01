package main

import "regexp"

func FixQuotes(text string) string {

    reDouble := regexp.MustCompile(`"\s*([^"]*?)\s*"`)
    text = reDouble.ReplaceAllString(text, `"$1"`)

    reSingle := regexp.MustCompile(`'\s*([^']*?)\s*'`)
    text = reSingle.ReplaceAllString(text, `'$1'`)

    return text
}