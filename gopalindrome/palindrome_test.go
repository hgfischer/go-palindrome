package gopalindrome

import (
	"testing"
)

var (
	palindromes = []string{
		"rotor!",
		"sixaxis",
		"@sixa$xis!",
		"Civic",
		"radar",
		"leVel",
		"noon",
		"Kayak",
		"reviver",
		"@racecar",
		"redder",
		"madam?!",
		"Malayalam",
		"refer",
		"l3v3l",
		"Eva, can I stab bats in a cave?",
		"Mr. Owl ate my metal worm",
		"Was it a car or a cat I saw?",
		"A nut for a jar of tuna",
		"Do geese see God?",
		"Ma is as selfless as I am",
		"On a clover, if alive erupts a vast pure evil, a fire volcano.",
		"Dammit, I'm mad!",
		"A Toyota's a Toyota",
		"Go hang a salami, I'm a lasagna hog",
		"A Santa lived as a devil at NASA",
	}

	nonPalindromes = []string{
		"",
		"fiveaxis",
		"5ixa5is",
		"Civik",
		"redar",
		"l3v31",
		"palindrome",
		"soon",
	}
)

func TestLowercase(t *testing.T) {
	if lowercase('A') != 'a' {
		t.Fatal("Lowercase failed")
	}
	if lowercase('Z') != 'z' {
		t.Fatal("Lowercase failed")
	}
}

func TestPalindromeA(t *testing.T) {
	for _, w := range palindromes {
		if !PalindromeA(w) {
			t.Fatal(w)
		}
	}
	for _, w := range nonPalindromes {
		if PalindromeA(w) {
			t.Fatal(w)
		}
	}
}

func BenchmarkPalindromeA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, w := range palindromes {
			PalindromeA(w)
		}
		for _, w := range nonPalindromes {
			PalindromeA(w)
		}
	}
}

func TestPalindromeB(t *testing.T) {
	for _, w := range palindromes {
		if !PalindromeB(w) {
			t.Fatal(w)
		}
	}
	for _, w := range nonPalindromes {
		if PalindromeB(w) {
			t.Fatal(w)
		}
	}
}

func BenchmarkPalindromeB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, w := range palindromes {
			PalindromeB(w)
		}
		for _, w := range nonPalindromes {
			PalindromeB(w)
		}
	}
}
