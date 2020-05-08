package _strings

import (
	"strings"
	"unicode/utf8"
)

func StringLeft (str string) []string {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return []string{"",""}
	}
	str = strings.ToLower(str)

	var pos = 0
	for k := 0; k < len(str); k++ {
		n, _ := utf8.DecodeRuneInString(string(str[k]))
		if n <= 47 || n >= 58 {
			pos++
		} else {
			break
		}
	}
	var symbol, number string
	if pos == 0 {
		symbol = str[pos:]
		number = ""
	} else {
		symbol = str[pos:]
		number = str[:pos]
	}
	return []string{number,symbol}
}
func StringRight (str string) []string {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return []string{"",""}
	}
	str = strings.ToLower(str)

	var pos = len(str)
	for k := len(str) - 1; k >= 0; k-- {
		n, _ := utf8.DecodeRuneInString(string(str[k]))
		if n <= 47 || n >= 58 {
			pos--
		} else {
			break
		}
	}
	var symbol, number string
	if pos == len(str) {
		symbol = str[:pos]
		number = ""
	} else {
		symbol = str[:pos]
		number = str[pos:]
	}
	return []string{number,symbol}
}
