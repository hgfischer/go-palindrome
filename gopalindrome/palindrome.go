// Package to compare the performance of different solutions for the same problem.
// Problem: given a non-empty string, return true if the string is a palindrome. Punctuation and spaces
// should be ignored.
package gopalindrome

import (
	"regexp"
)

var (
	regex = regexp.MustCompile("[^a-zA-Z]")
)

// Quick and dirty palindrome solution using regex to cleanup string
func PalindromeA(s string) bool {
	s = regex.ReplaceAllString(s, "")
	l := len(s)
	if l == 0 {
		return false
	}
	for i := 0; i < l/2; i++ {
		if lowercase(s[i]) != lowercase(s[l-(i+1)]) {
			return false
		}
	}
	return true
}

// Optimized palindrome solution with almost O(n) complexity
func PalindromeB(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	last := len(s) - 1
	for i := 0; i <= l/2; i++ {
		if !isAlpha(s[i]) {
			continue
		}
		for j := last; j >= l/2; j-- {
			last -= 1
			if !isAlpha(s[j]) {
				continue
			}

			if lowercase(s[i]) == lowercase(s[j]) {
				break
			} else {
				return false
			}
		}
	}
	return true
}

func isAlpha(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}

func lowercase(char byte) byte {
	if 'A' <= char && char <= 'Z' {
		char += 'a' - 'A'
	}
	return char
}
