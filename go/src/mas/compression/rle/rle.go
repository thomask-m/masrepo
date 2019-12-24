package rle

import (
	"strconv"
)

func Encode(s string) string {
	ret := ""
	prev := ""
	count := 0
	for _, curr := range s {
		if string(curr) != prev && count != 0 {
			ret += strconv.Itoa(count) + prev
			count = 0
		} else {
			prev = string(curr)
			count++
		}
	}
	if count != 0 {
		ret += strconv.Itoa(count) + prev
	}
	return ret
}

func Decode(s string) string {
	return ""
}
