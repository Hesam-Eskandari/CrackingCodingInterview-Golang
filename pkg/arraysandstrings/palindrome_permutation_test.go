package arraysandstrings

import "testing"

func (p *palindromePermutation) assertTrue(t *testing.T, output bool) {
	if !output {
		t.Errorf("expected true returned false, string: \"%s\", ignore_case: %v", p.string, p.ignoreCase)
	}
}

func (p *palindromePermutation) assertFalse(t *testing.T, output bool) {
	if output {
		t.Errorf("expected false returned true, string: \"%s\", ignore_case: %v", p.string, p.ignoreCase)
	}
}

func TestPalindromePermutation(t *testing.T) {
	inputTrue := []*palindromePermutation{
		NewPalindromePermutation("Taco  cat", true),
		NewPalindromePermutation("tacocat", false),
		NewPalindromePermutation("123TaCo C A T_ 123_", true),
		NewPalindromePermutation("", true),
		NewPalindromePermutation("", false),
		NewPalindromePermutation(" ", true),
		NewPalindromePermutation(" ", false),
		NewPalindromePermutation("!123444a ", false),
		NewPalindromePermutation("taco cat", false),
	}
	inputFalse := []*palindromePermutation{
		NewPalindromePermutation("Tacocat", false),
		NewPalindromePermutation("Taco cat", false),
		NewPalindromePermutation("Tacot cat", true),
	}
	for _, input := range inputTrue {
		t.Run("should return true", func(t *testing.T) {
			input.assertTrue(t, input.isPalindromePermute())
		})
	}
	for _, input := range inputFalse {
		t.Run("should return false", func(t *testing.T) {
			input.assertFalse(t, input.isPalindromePermute())
		})
	}
}
