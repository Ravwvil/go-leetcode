package main

import "fmt"

func main() {
	var s string
	var t string
	fmt.Scanf("%s \n %s", &s, &t)
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	var s_alphabet [26]int
	var t_alphabet [26]int

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		s_alphabet[s[i]-'a']++
		t_alphabet[t[i]-'a']++
	}

	if s_alphabet == t_alphabet {
		return true
	}

	return false
}
