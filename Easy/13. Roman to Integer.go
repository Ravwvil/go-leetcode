package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && symbols[s[i]] < symbols[s[i+1]] {
			result -= symbols[s[i]]
		} else {
			result += symbols[s[i]]
		}
	}
	return result
}
