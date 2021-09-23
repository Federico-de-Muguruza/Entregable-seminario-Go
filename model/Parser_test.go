package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parse
		Success bool   // parser resul
		Type    string // the input typ
		Value   string // the input valu
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX05ABC5E", false, "TX", "ABC5E", 5},
		{"NN04000A", false, "NN", "000A", 4},
		{"TX06ABCDE", false, "TX", "ABCDE", 5},
		{"TX0ZABCDE", false, "TX", "ABCDE", 5},
		{"ZA03ABE", false, "ZA", "ABE", 3},
	}

	for _, c := range cases {
		_, err := Parser(c.Input)
		assert.Equal(t, err == nil, c.Success)
	}
}
