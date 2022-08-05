package tests

import (
	"strings"
	"task/deeSeeComics/cryptography"
	"testing"
)

type encodingTest struct {
	input    string
	key      int
	expected string
}

var tests = []encodingTest{
	{"clark", 5, "hqfwp"},
	{"blossom", 3, "eorvvrp"},
}

func TestDeeSeeChiffreShouldReturnValidEncoding(t *testing.T) {
	for _, test := range tests {
		if output := strings.Map(func(r rune) rune {
			return cryptography.DeeSeeChiffre(r, test.key)
		}, test.input); output != test.expected {
			t.Errorf("expected value was %v, but got %v", test.expected, output)
		}
	}
}
