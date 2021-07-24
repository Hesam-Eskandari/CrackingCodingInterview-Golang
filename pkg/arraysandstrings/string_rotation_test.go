package arraysandstrings

import "testing"

func (sr *stringRotation) assertTrue(t *testing.T, out bool) {
	if !out {
		t.Errorf("expected true, returned %v, first string: %s, second string: %s",
			out, sr.firstString, sr.secondString)
	}
}

func (sr *stringRotation) assertFalse(t *testing.T, out bool) {
	if out {
		t.Errorf("expected false, returned %v, first string: %s, second string: %s",
			out, sr.firstString, sr.secondString)
	}
}

func TestStringRotation(t *testing.T) {
	inputs_true := []*stringRotation{
		NewStringRotation("", ""),
		NewStringRotation("ab", "ba"),
		NewStringRotation("aAAa", "aaAA"),
		NewStringRotation("abcdefg", "cdefgab"),
		NewStringRotation("abcdefg", "abcdefg"),
	}
	inputs_false := []*stringRotation{
		NewStringRotation(" ", "."),
		NewStringRotation("abcdefg", "abcdef"),
		NewStringRotation("a b c d", "a b cd "),
		NewStringRotation("abcdefg", "cDefgab"),
	}
	for _, input := range inputs_true {
		t.Run("should return true", func(t *testing.T) {
			input.assertTrue(t, input.areRotations())
		})
	}
	for _, input := range inputs_false {
		t.Run("should return false", func(t *testing.T) {
			input.assertFalse(t, input.areRotations())
		})
	}
}
