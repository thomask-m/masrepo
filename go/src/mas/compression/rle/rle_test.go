package rle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := map[string]string{
		"aaaaa": "5a",
	}
	for input, expected := range tests {
		assert.Equal(t, expected, Encode(input))
	}
}
