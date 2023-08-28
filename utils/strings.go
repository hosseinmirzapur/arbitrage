package utils

import (
	"strings"
)

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Hyphenize(base, quote string) string {
	return base + "-" + quote
}

func DeHyphenize(s string) string {
	return strings.Replace(s, "-", "", -1)
}
