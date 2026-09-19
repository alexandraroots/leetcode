package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isValidPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return r
	}, s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	strs := "A man, a plan, a canal: Panama"
	fmt.Println(isValidPalindrome(strs))

}
