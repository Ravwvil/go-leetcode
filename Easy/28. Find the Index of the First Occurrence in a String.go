package main

import "fmt"

func main() {
	var haystack string
	var needle string
	fmt.Println("Enter 2 strings separated by space:")
	fmt.Scan(&haystack, &needle)
	result := strStr(haystack, needle)
	if result != -1 {
		fmt.Println("The index of the first occurrence is:", result)
	} else {
		fmt.Println("Needle not found in haystack")
	}
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	if needle == "" {
		return 0
	}
	return -1
}
