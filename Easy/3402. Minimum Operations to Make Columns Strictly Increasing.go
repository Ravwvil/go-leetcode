package easy

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := [][]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rowStrs := strings.Fields(line)
		row := make([]int, len(rowStrs))
		for j, numStr := range rowStrs {
			row[j], _ = strconv.Atoi(numStr)
		}
		grid = append(grid, row)
	}
	println(minimumOperations(grid))
}

func minimumOperations(grid [][]int) int {
	counter := 0

	for j := 0; j < len(grid[0]); j++ {
		for i := 1; i < len(grid); i++ {
			if grid[i][j] <= grid[i-1][j] {
				counter += grid[i-1][j] - grid[i][j] + 1
				grid[i][j] = grid[i-1][j] + 1
			}
		}
	}
	return counter
}
