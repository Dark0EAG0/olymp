package main

func lengthOfLongestSubstring(s string) int {
	l := 0
	for i := 0; i < len(s); i++ {
		b := fn(s[i:])
		if b > l {
			l = b
		}
	}
	return l
}

func fn(s string) int {
	a := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if is, _ := a[s[i]]; is {
			return i
		} else {
			a[s[i]] = true
		}
	}
	return len(s)
}
