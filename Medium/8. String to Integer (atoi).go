package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i >= len(s) {
		return 0
	}

	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	result := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && digit > 7) {
			if sign == 1 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}
