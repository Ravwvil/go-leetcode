package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)
	fmt.Println(reverse(x))
}

// Math solution
func reverse(x int) int {
	reverse := 0
	for x != 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse
}

// String solution

/*
func reverse(x int) int {
	str := strconv.Itoa(x)
	var returnStr string
	for i := len(str) - 1; i >= 0; i-- {
		returnStr += string(str[i])
	}
	if x < 0 {
		returnStr = "-" + returnStr[:len(returnStr)-1]
	}
	reverse, _ := strconv.Atoi(returnStr)
	if reverse > 1<<31-1 || reverse < -1<<31 {
		return 0
	}
	return reverse
}
*/
