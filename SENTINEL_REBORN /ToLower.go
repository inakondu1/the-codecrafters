package main

import (
	"strings"
)

func ToLower(s string) string {
	w := strings.Fields(s)
	for i := 1; i < len(w); i++ {
		if w[i] == "(low)" {
			w[i-1] = strings.ToLower(w[i-1])
			w = append(w[:i], w[i+1:]...)
		}
	}
	return strings.Join(w, " ")
}
