package viper

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

func checkKey(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("%s key is not set", key))
	}
}

func panicIfErrorForKey(err error, key string) {
	if err != nil {
		panic(fmt.Sprintf("Could not parse key: %s, Error: %s", key, err))
	}
}

func GetStringOrPanic(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func GetIntOrPanic(key string) int {
	v, err := strconv.Atoi(GetStringOrPanic(key))
	panicIfErrorForKey(err, key)
	return v
}

func GetInt64OrPanic(key string) int64 {
	v, err := strconv.ParseInt(GetStringOrPanic(key), 10, 64)
	panicIfErrorForKey(err, key)
	return v
}

func GetBool(key string) bool {
	v, err := strconv.ParseBool(GetString(key))
	if err != nil {
		return false
	}
	return v
}

func GetBoolOrPanic(key string) bool {
	v, err := strconv.ParseBool(GetString(key))
	panicIfErrorForKey(err, key)
	return v
}

func GetString(key string) string {
	return viper.GetString(key)
}

func SplitStringOrPanic(key string, delimiter string) []string {
	checkKey(key)
	slice := splitAndTrim(viper.GetString(key), delimiter)
	return slice
}

func splitAndTrim(a string, delimiter string) []string {
	slice := strings.Split(a, delimiter)
	for key, value := range slice {
		slice[key] = strings.TrimSpace(value)
	}
	return slice
}
