package easy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nums []int
	var target int

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for _, s := range strings.Fields(input) {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}

	fmt.Scan(&target)

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	pairs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := pairs[nums[i]]; ok {
			return []int{pairs[nums[i]], i}
		}
		pairs[target-nums[i]] = i
	}
	return []int{0, 0}
}
