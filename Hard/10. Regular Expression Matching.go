package main

import "fmt"

func main() {
	var s string
	var p string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &p)
	fmt.Println(isMatch(s, p))
}

// DP solution
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true

	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' && j >= 2 {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				if j >= 2 {
					dp[i][j] = dp[i][j-2]
				}

				if j >= 2 && (p[j-2] == '.' || p[j-2] == s[i-1]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[len(s)][len(p)]
}

/*
Recursive Solution
func isMatch(s string, p string) bool {
	memo := make(map[string]bool)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		key := fmt.Sprintf("%d,%d", i, j)
		if result, ok := memo[key]; ok {
			return result
		}

		if j == len(p) {
			return i == len(s)
		}

		firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')

		var result bool

		if j+1 < len(p) && p[j+1] == '*' {
			result = dp(i, j+2) || (firstMatch && dp(i+1, j))
		} else {
			result = firstMatch && dp(i+1, j+1)
		}
		memo[key] = result
		return result
	}

	return dp(0, 0)
}
*/
