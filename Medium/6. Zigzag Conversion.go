package main

import "fmt"

func main() {
	var s string
	var numRows int
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &numRows)
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	result := make([]byte, 0, len(s))
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += cycleLen {
			result = append(result, s[j])

			secondIdx := j + cycleLen - 2*i
			if i > 0 && i < numRows-1 && secondIdx < len(s) {
				result = append(result, s[secondIdx])
			}
		}
	}

	return string(result)
}
