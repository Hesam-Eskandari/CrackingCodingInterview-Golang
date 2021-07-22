package arraysandstrings

import "testing"

func (u *url) assertEqual(t *testing.T, expected string) {
	if u.string != expected {
		t.Errorf("expected %s, returned %s", expected, u.string)
	}
}

func TestUrlify(t *testing.T) {
	inputs := []url{{string: "Mr John Smith    ", length: 13},
		{string: " Mr John Smith      ", length: 14},
		{string: " Mr  John Smith           ", length: 16}}
	expects := []string{"Mr%20John%20Smith", "%20Mr%20John%20Smith", "%20Mr%20%20John%20Smith%20"}
	for index, input := range inputs {
		t.Run("test urlify", func(t *testing.T) {
			u := NewURLify(input.string, input.length)
			u.Urlify()
			u.assertEqual(t, expects[index])
		})
	}
}
