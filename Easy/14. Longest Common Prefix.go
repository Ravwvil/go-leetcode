package easy

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var strs []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		strs = append(strs, line)
	}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j == len(strs[i]) || prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}

	return prefix
}
