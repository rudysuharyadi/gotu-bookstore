package utils

import (
	"strconv"
	"strings"
)

func SplitAndTrim(a string, delimiter string) []string {
	slice := strings.Split(a, delimiter)
	for key, value := range slice {
		slice[key] = strings.TrimSpace(value)
	}
	return slice
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SliceToMap(slice []string) map[string]string {
	elementMap := make(map[string]string)
	for _, s := range slice {
		elementMap[s] = s
	}
	return elementMap
}

func SliceInSlice(a []string, list []string) bool {
	elementMap := SliceToMap(list)
	for _, b := range a {
		_, ok := elementMap[b]
		if !ok {
			return false
		}
	}
	return true
}

func StringToInt64(a string) int64 {
	if n, err := strconv.ParseInt(a, 10, 64); err == nil {
		return n
	}
	return 0
}

func FloatToString(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func Int64ToString(a int64) string {
	return strconv.FormatInt(a, 10)
}

func IntToString(a int) string {
	return strconv.Itoa(a)
}

func StringToInt(a string) int {
	if n, err := strconv.Atoi(a); err == nil {
		return n
	}
	return 0
}

func StringToFloat64(a string) float64 {
	if n, err := strconv.ParseFloat(a, 64); err == nil {
		return n
	}
	return 0
}

func BooleanToString(a bool) string {
	return strconv.FormatBool(a)
}

// GetStringInBetween Returns empty string if no start string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	e = s + e
	return str[s:e]
}

func Substring(str string, start int, end int) (result string) {
	if start < 0 {
		start = 0
	}
	if end < 0 {
		return
	}
	if start > len(str) {
		return
	}
	if end > len(str) {
		end = len(str)
	}
	if start > end {
		return
	}
	return str[start:end]
}
