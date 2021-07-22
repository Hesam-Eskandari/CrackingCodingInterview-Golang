package arraysandstrings

import "testing"

type testCaseStringCompression struct {
	message           string
	stringCompression *stringCompression
	expected          string
}

func (sc *testCaseStringCompression) assertEqual(t *testing.T, out string, expected string) {
	if out != expected {
		t.Errorf("did not return the correct value. expected: \"%s\", returned; \"%s\"", expected, out)
	}
}

func TestStringCompression(t *testing.T) {
	inputs := []*testCaseStringCompression{
		{"empty string", NewStringCompression(""), ""},
		{"string of length 1", NewStringCompression("a"), "a"},
		{"string of length 2", NewStringCompression("aa"), "aa"},
		{"three of a kind", NewStringCompression("aaa"), "a3"},
		{"long challenging string", NewStringCompression("aabaAbbbaaaa"), "a2b1a1A1b3a4"},
		{"case sensitivity", NewStringCompression("aaAAaaa"), "a2A2a3"},
		{"11 of a kind", NewStringCompression("aaaaaaaaaaa"), "a11"},
	}
	for _, input := range inputs {
		t.Run(input.message, func(t *testing.T) {
			input.assertEqual(t, input.stringCompression.compressString(), input.expected)
		})
	}
}
