package utils

import "strings"

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s // Return the string as is if it's empty
	}
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}
	return strings.Join(words, " ")
}
