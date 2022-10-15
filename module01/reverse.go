package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"

// The next three functions have a problem handling letters
// greater than one byte, last one handles this case
func Reverse(word string) string {
	var r string

	for i := len(word) - 1; i >= 0; i-- {
		r = r + string(word[i])
	}

	return r
}

func ReverseEff(word string) string {
	// Builder is struct with methods for efficient string handling
	var sb strings.Builder

	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i])
	}

	return sb.String()
}

func ReverseUp(word string) string {
	var r string

	for i := 0; i < len(word); i++ {
		r = string(word[i]) + r
	}

	return r
}

func ReverseRunes(word string) string {
	var r string

	// Iterating over strings like this gives an
	// individual rune for letter instead of
	// indexing the bytes
	for _, letter := range word {
		r = string(letter) + r
	}

	return r
}
