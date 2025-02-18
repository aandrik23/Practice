package main

func isValidCamelCase(s string) bool {
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return false
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return false
		}
	}

	for _, char := range s {
		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z') {
			return false
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	if !isValidCamelCase(s) {
		return s
	}

	var result string

	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i != 0 {
				result += "_"
			}
		}
		result += string(char)
	}
	return result
}
