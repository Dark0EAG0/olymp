package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := ""
	for i := 0; i < len(word1); i++ {
		s1 += word1[i]
	}
	s2 := ""
	for i := 0; i < len(word2); i++ {
		s2 += word2[i]
	}
	if s1 == s2 {
		return true
	}
	return false
}
