package main

import "unicode"

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) {
			count++
		}
	}
	return count
}
