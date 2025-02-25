package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	substring := s[:1]
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)
		if len(odd) > len(substring) {
			substring = odd
		}
		if len(even) > len(substring) {
			substring = even
		}
	}
	return substring
}

func expand(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
