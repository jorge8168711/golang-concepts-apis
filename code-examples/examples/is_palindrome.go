package examples

import "strings"

func IsPalindrome(text string) bool {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		// text[i] return the ascii code
		// and string convert the ascii code to string
		textReverse += string(text[i])
	}

	return strings.EqualFold(strings.ToLower(text), strings.ToLower(textReverse))
}
