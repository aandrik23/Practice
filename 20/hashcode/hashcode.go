package main

func HashCode(dec string) string {
	result := ""
	strLen := len(dec)

	for _, char := range dec {
		hasVal := (int(char) + strLen) % 127

		if hasVal < 32 {
			hasVal += 33
		}
		result += string(rune(hasVal))
	}
	return result
}
