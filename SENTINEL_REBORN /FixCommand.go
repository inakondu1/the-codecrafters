package main

import (
	"strconv"
	"strings"
	"unicode"
)

func capitalizeWord(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	for i := 1; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}

func edit(s string) string {
	w := strings.Fields(s)

	for i := 1; i < len(w); i++ {
		if strings.HasPrefix(w[i], "(") {
			key := w[i]
			usedNext := false
			if !strings.HasSuffix(key, ")") && i+1 < len(w) {
				key = key + " " + w[i+1]
				usedNext = true
			}
			key = strings.TrimSpace(strings.Trim(key, "()"))
			parts := strings.Split(key, ",")
			command := strings.TrimSpace(parts[0])
			count := 1
			if len(parts) == 2 {
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}
			for j := 1; j <= count && i-j >= 0; j++ {
				switch command {
				case "up":
					w[i-j] = strings.ToUpper(w[i-j])
				case "low":
					w[i-j] = strings.ToLower(w[i-j])
				case "cap":
					w[i-j] = capitalizeWord(w[i-j])
				}
			}
			if usedNext {
				w = append(w[:i], w[i+2:]...)
			} else {
				w = append(w[:i], w[i+1:]...)
			}
			i--
		}
	}
	return strings.Join(w, " ")
}
