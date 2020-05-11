package _string

import (
	"strings"
	"unicode/utf8"
)

//получить числовую подстроку слева => num,symbols
func NumberLeft(str string) (string, string) {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return "", ""
	}
	str = strings.ToLower(str)

	var pos = 0
	for k := 0; k < len(str); k++ {
		n, _ := utf8.DecodeRuneInString(string(str[k]))
		if n > 47 && n < 58 {
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
	return number, symbol
}

//получить числовую подстроку справа => num,symbols
func NumberRight(str string) (string, string) {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return "", ""
	}
	str = strings.ToLower(str)

	var pos = len(str)
	for k := len(str) - 1; k >= 0; k-- {
		n, _ := utf8.DecodeRuneInString(string(str[k]))
		if n > 47 && n < 58 {
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
	return number, symbol
}
