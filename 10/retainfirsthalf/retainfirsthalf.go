package main

func RetainFirstHalf(str string) string {
	length := len(str)

	if length == 0 {
		return ""
	}

	if length == 1 {
		return str
	}
	halfLength := length / 2

	return str[:halfLength]
}
