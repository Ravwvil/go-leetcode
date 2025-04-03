package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	rowStrs := strings.Fields(line)
	for i := range rowStrs {
		val, _ := strconv.Atoi(rowStrs[i])
		nums = append(nums, val)
	}
	fmt.Println(containsDuplicate(nums))
}

// Complexity: O(n^2) worst case, O(n) best case, runtime 0ms on LeetCode
func containsDuplicate(nums []int) bool {
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == nums[i+1] {
				return true
			}
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				isSorted = false
			}
		}
	}
	return false
}

// Complexity: O(n), runtime 13ms on LeetCode
//func containsDuplicate(nums []int) bool {
//	m := make(map[int]bool)
//	for _, num := range nums {
//		if m[num] {
//			return true
//		}
//		m[num] = true
//	}
//	return false
//}
