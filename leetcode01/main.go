package main

import (
	"log"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re, err := regexp.Compile(`[^a-zA-Z0-9]`)
	if err != nil {
		log.Fatal(err)
	}
	s = strings.ReplaceAll(strings.ToLower(re.ReplaceAllString(s, "")), " ", "")
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
