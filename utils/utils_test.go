package utils

import (
	"testing"
)

func TestGetCmdFromPrompt(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: "Hello wOrld", expected: []string{"hello", "world"}},
		{input: "Lorem ipSum", expected: []string{"lorem", "ipsum"}},
	}

	for _, cs := range cases {
		slicedPrompt, _ := GetCmdFromPrompt(cs.input)
		for i := range slicedPrompt {
			actual := slicedPrompt[i]
			expected := cs.expected[i]
			if actual != expected {
				t.Errorf("%s doesn't equals %s", actual, expected)
			}
		}
	}
}

func TestGetOffsetFromUrl(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{input: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20", expected: 20},
		{input: "https://pokeapi.co/api/v2/location-area/?limit=20", expected: 0},
		{input: "", expected: 0},
	}

	for _, cs := range cases {
		offset := GetOffsetFromUrl(&cs.input)
		expected := cs.expected
		if offset != expected {
			t.Errorf("Actual %d doesn't equals %d", offset, expected)
		}
	}
}
