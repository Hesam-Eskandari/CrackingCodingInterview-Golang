package arraysandstrings

import "testing"

func (oa *oneAway) assertTrue(t *testing.T, out bool) {
	if !out {
		t.Errorf("expected true, returned false. first string: \"%s\", second string: \"%s\"", oa.firstString, oa.secondString)
	}
}

func (oa *oneAway) assertFalse(t *testing.T, out bool) {
	if out {
		t.Errorf("expected false, returned true. first string: \"%s\", second string: \"%s\"", oa.firstString, oa.secondString)
	}
}

func TestOnwAway(t *testing.T) {
	inputsTrue := []*oneAway{
		NewOnwAway("", ""),
		NewOnwAway("", " "),
		NewOnwAway("a", "ab"),
	}
	inputsFalse := []*oneAway{
		NewOnwAway("abc", "a"),
		NewOnwAway("abc", "acb"),
		NewOnwAway(" a", "a "),
	}
	for _, input := range inputsTrue {
		t.Run("should return true", func(t *testing.T) {
			input.assertTrue(t, input.isOneAway())
		})
	}
	for _, input := range inputsFalse {
		t.Run("should return false", func(t *testing.T) {
			input.assertFalse(t, input.isOneAway())
		})
	}
}
