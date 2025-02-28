package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	start := 0
	maxLength := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			start = max(start, m[s[i]]+1)
		}
		m[s[i]] = i
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
