package main

import "testing"

func TestCleanInput(t *testing.T) {
cases := []struct {
	input string
	expected []string
}{
	{
		input:	"hello world",
		expected: []string{"hello", "world"},
	},
	{
		input: "professor oak",
		expected: []string{"professor", "oak"},
	},
	{
		input: "pewter city",
		expected: []string{"pewter", "city"},
	},
}
for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected) {
		t.Errorf("Word count incorrect. Expected %d, Actual %d", len(c.expected), len(actual))
	}
	for i := range actual {
		word := actual[i]
		if word != c.expected[i] {
			t.Errorf("Words do not match. Expected %s, Actual %s", c.expected[i], word)
		}
	}
}
}