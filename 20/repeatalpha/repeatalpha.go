package main

import "strings"

func RepeatAlpha(s string) string {
	var result strings.Builder

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			var repeatCount int
			if r >= 'a' && r <= 'z' {
				repeatCount = int(r - 'a' + 1)
			} else {
				repeatCount = int(r - 'A' + 1)
			}
			for i := 0; i < repeatCount; i++ {
				result.WriteRune(r)
			}
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
