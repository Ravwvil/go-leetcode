package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var nums []int
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	for _, str := range strings.Split(input, " ") {
		var num int
		fmt.Sscanf(str, "%d", &num)
		nums = append(nums, num)
	}
	fmt.Println(canBeIncreasing(nums))
}

func canBeIncreasing(nums []int) bool {
	var n = len(nums)
	var changed bool

	for i := 0; i < n-1; i++ {
		if nums[i] >= nums[i+1] {
			if changed {
				return false
			}
			changed = true
			if i > 0 && nums[i-1] >= nums[i+1] {
				nums[i+1] = nums[i]
			}
		}
	}

	return true
}
