package utils

import (
	"log"
	"strconv"
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

func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func StringToFloat(s string) float64 {
	floatNum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}

	return floatNum
}
