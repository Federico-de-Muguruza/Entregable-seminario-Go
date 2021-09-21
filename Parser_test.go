package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"tudai2021.com/model"
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
		// {"TX06ABCDE", false, "", "", },
		{"NN04000A", false, "", "", 0},
	}

	for _, c := range cases {
		_, err := model.Parser(c.Input)
		assert.Equal(t, err, nil, c.Success)
	}
}
