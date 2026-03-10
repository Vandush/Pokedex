package functions

import (
	"testing"
) 
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "wabajack",
			expected: []string{"wabajack"},
		},
		{
			input: "The Great Bandito!",
			expected: []string{"the", "great", "bandito!"},
		},
	}
	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected: %v, Actual: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected: %s, Got: %s", expectedWord, word)
			}
		}
	}
}
