package rle

func Encode(s string) string {
	ret := ""
	prev := ""
	count := 0
	for _, curr := range s {
		if string(curr) != prev && count != 0 {
			ret += string(count) + prev
			count = 0
			continue
		}
		count++
	}
	return ret
}

func Decode(s string) string {
	return ""
}
