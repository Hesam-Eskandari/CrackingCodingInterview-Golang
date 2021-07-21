package arraysandstrings

import "testing"

func (s *uniqueString) assertTrue(t *testing.T, output bool) {
	if !output {
		t.Errorf("expected true, got %v", output)
	}
}

func (s *uniqueString) assertFalse(t *testing.T, output bool) {
	if output {
		t.Errorf("expected false, got %v", output)
	}
}

func TestIsUnique(t *testing.T) {
	inputsTrue := []string{"123qweasdzxc ", "12a", "_123 "}
	inputsFalse := []string{"123qweasdzxc  ", "_12a1", "_123?? "}
	for _, input := range inputsTrue {
		t.Run("should return true", func(t *testing.T) {
			uString := NewUniqueString(input)
			output := uString.IsUnique()
			uString.assertTrue(t, output)
		})
	}
	for _, input := range inputsFalse {
		t.Run("should return false", func(t *testing.T) {
			uString := NewUniqueString(input)
			output := uString.IsUnique()
			uString.assertFalse(t, output)
		})
	}

}
