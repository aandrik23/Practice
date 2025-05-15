package main

import "strings"

func LastWord(s string) string {
	words := strings.Fields(s)
	for len(words) == 0 {
		return "\n"
	}
	return words[len(words)-1] + "\n"
}
