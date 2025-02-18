package main

func FirstWord(s string) string {
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
		word += string(s[i])
	}
	return word + "\n"
}
