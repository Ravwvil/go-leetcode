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
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (r == ')' && top != '(') || (r == ']' && top != '[') || (r == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}
