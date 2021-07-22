package arraysandstrings

import "testing"

func (cp *checkPermutation) assertTrue(t *testing.T, output bool) {
	if !output {
		t.Errorf("expected true, returned false: %s and %s do not seem to be eachother's permutation",
			cp.firstString, cp.secondString)
	}

}

func (cp *checkPermutation) assertFalse(t *testing.T, output bool) {
	if output {
		t.Errorf("expected false, returned true: \"%s\" and \"%s\" do seem to be eachother's permutation",
			cp.firstString, cp.secondString)
	}

}

func TestCheckPermutation(t *testing.T) {
	inputsTrue := [][]string{{"taco act", "cat octa"}, {"tt  aa123", "1t a2t a3"}, {"", ""}, {" ", " "}}
	inputsFalse := [][]string{{"Taco act", "taco act"}, {"", " "}, {"string", "rings"}, {"tacoact", "taco act"}}
	for _, list := range inputsTrue {
		t.Run("must return true", func(t *testing.T) {
			cp := NewCheckPermutation(list[0], list[1])
			output := cp.arePermute()
			cp.assertTrue(t, output)
		})
	}
	for _, list := range inputsFalse {
		t.Run("must return false", func(t *testing.T) {
			cp := NewCheckPermutation(list[0], list[1])
			output := cp.arePermute()
			cp.assertFalse(t, output)
		})
	}
}
