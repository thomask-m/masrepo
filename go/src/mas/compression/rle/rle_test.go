package rle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests map[string]string = map[string]string{
	"aaaaa":                     "5a",
	"a":                         "1a",
	"aabbc":                     "2a2b1c",
	"asdfasdf":                  "1a1s1d1f1a1s1d1f",
	"uuuhhuuu":                  "3u2h3u",
	"###$$%":                    "3#2$1%",
	"uuuuuuuuuuuiiiiiiiiiiiooo": "11u11i3o",
}

func TestEncode(t *testing.T) {
	for input, expected := range tests {
		assert.Equal(t, expected, Encode(input))
	}
}

func TestDecode(t *testing.T) {
	for expected, input := range tests {
		assert.Equal(t, expected, Decode(input))
	}
}
