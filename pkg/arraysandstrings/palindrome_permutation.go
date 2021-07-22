package arraysandstrings

import (
	"strings"
)

/*
Cracking the coding interview, 6th edition, page 90 ((Modified Assumptions))
Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome.
Version one: case sensitive and whitespace significant (ignoreCase = false) => letters only
Version two: case insensitive and whitespace insignificant (ignoreCase = true) => any byte
*/

type palindromePermutation struct {
	string     string
	ignoreCase bool
	minLower   byte
	minUpper   byte
	maxLower   byte
	maxUpper   byte
}

func NewPalindromePermutation(string string, ignoreCase bool) *palindromePermutation {
	var minLower, minUpper, maxLower, maxUpper byte
	if ignoreCase {
		string = strings.ToLower(string)
		minLower = 0
		minUpper = 0
		maxLower = 255
		maxUpper = 255
	} else {
		minLower = 65
		minUpper = 97
		maxLower = 90
		maxUpper = 122
	}
	return &palindromePermutation{
		string:     string,
		ignoreCase: ignoreCase,
		minLower:   minLower,
		minUpper:   minUpper,
		maxLower:   maxLower,
		maxUpper:   maxUpper,
	}
}

func (p *palindromePermutation) isPalindromePermute() bool {
	split := strings.Split(p.string, "")
	hmap := make(map[string]int)
	for _, str := range split {
		bytePosition := []byte(str)[0]
		if (bytePosition >= p.minLower && bytePosition <= p.maxLower) ||
			(bytePosition >= p.minUpper && bytePosition <= p.maxUpper) {
			if _, ok := hmap[str]; ok {
				hmap[str] += 1
			} else {
				hmap[str] = 1
			}
		}
	}
	countOdd := 0
	for _, value := range hmap {
		if value%2 != 0 {
			countOdd += 1
		}
	}
	return countOdd <= 1
}
