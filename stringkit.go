// Package stringkit is a small, dependency-free collection of string helpers.
//
// It exists as a deliberately tiny, well-tested target for exercising
// automated issue-to-PR workflows: each function is simple enough that a new
// feature request maps cleanly onto a single change with an obvious test.
package stringkit

import (
	"strings"
	"unicode"
)

// Reverse returns s with its characters in reverse order. It operates on runes,
// so multibyte characters (e.g. "áé", "日本語") are reversed correctly rather
// than split mid-byte.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same forwards and backwards, ignoring
// case and any character that is not a letter or digit. The empty string is
// considered a palindrome.
func IsPalindrome(s string) bool {
	var cleaned []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}
	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}
	return true
}

// WordCount returns the number of whitespace-separated words in s. Leading,
// trailing, and repeated whitespace are ignored, so " hello   world " counts
// as 2.
func WordCount(s string) int {
	return len(strings.Fields(s))
}
