package util

import (
	"strconv"
	"strings"
)

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StrIsEmpty(s string) bool {
	if len(s) <= 0 {
		return true
	}

	if len(strings.TrimSpace(s)) <= 0 {
		return true
	}
	return false
}
