package rle

import (
	"strconv"
	"strings"
	"unicode"
)

func Encode(s string) string {
	if len(s) == 0 {
		return ""
	}

	ret := ""
	prev := ""
	count := 0

	for _, curr := range s {
		if string(curr) != prev && count != 0 {
			ret += strconv.Itoa(count) + prev
			count = 1
		} else {
			count++
		}
		prev = string(curr)
	}

	ret += strconv.Itoa(count) + prev
	return ret
}

func Decode(s string) string {
	ret := ""
	cs := strings.FieldsFunc(s, func(c rune) bool {
		return unicode.IsNumber(c)
	})

	nums := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsNumber(c)
	})

	for i, c := range cs {
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			return ""
		}
		ret += strings.Repeat(string(c), n)
	}
	return ret
}
