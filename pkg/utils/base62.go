package utils

import "strings"

const base62Chars = "Rk8jXnZWPq6wV9GDYJu2L7NxFTcyBghK0lQiHa4OzCbSdEfmMUs51Pp3vIrYt"

func ConvertBase62(input int64) string {
	if input == 0 {
		return "0"
	}

	var result strings.Builder
	for input > 0 {
		result.WriteByte(base62Chars[input%62])
		input /= 62
	}

	return ReverseString(result.String())
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
