package strutils

import (
	"strings"
	"unicode"
)

// ToUpperCase Returns the string changed with Uppercase
func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

// CapFirstLetter capitalise first letter of a string
func CapFirstLetter(s string) string {
	if len(s) < 1 { // if the string is empty
		return s
	}
	// trim the string
	t := strings.Trim(s, " ")
	// convert all letters to lower case
	t = strings.ToLower(t)
	res := []rune(t)
	// convert the first letter to uppercase
	res[0] = unicode.ToUpper(res[0])
	return string(res)
}
