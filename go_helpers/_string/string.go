package _string

import "unicode"

func IsLower(s string) bool {
	for _, r := range []rune(s) {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
