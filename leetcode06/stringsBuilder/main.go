package main

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var b1 strings.Builder
	var b2 strings.Builder

	for _, w := range word1 {
		b1.WriteString(w)
	}
	for _, w := range word2 {
		b2.WriteString(w)
	}
	return b1.String() == b2.String()
}
