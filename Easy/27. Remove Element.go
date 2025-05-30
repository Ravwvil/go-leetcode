package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := make([]int, 0)
	var number int
	var input string

	fmt.Println("Enter number that will be deleted:")
	fmt.Scan(&number)
	fmt.Println("Enter numbers separated by space:")
	fmt.Scanln(&input)

	parts := strings.Fields(input)

	for _, p := range parts {
		val, err := strconv.Atoi(p)
		if err != nil {
			fmt.Println(err)
			continue
		}
		nums = append(nums, val)
	}
	removeVal := removeElement(nums, number)
	fmt.Println(nums)
	fmt.Println(removeVal)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			i--
		}
	}
	return len(nums)
}
