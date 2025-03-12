package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums1 := []int{}
	nums2 := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	rowStrs := strings.Fields(line)
	for i := range rowStrs {
		val, _ := strconv.Atoi(rowStrs[i])
		nums1 = append(nums1, val)
	}
	scanner.Scan()
	line = scanner.Text()
	rowStrs = strings.Fields(line)
	for i := range rowStrs {
		val, _ := strconv.Atoi(rowStrs[i])
		nums2 = append(nums2, val)
	}
	println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	nums3 := make([]int, n+m)
	i, j := 0, 0
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			nums3[i+j] = nums1[i]
			i++
		} else {
			nums3[i+j] = nums2[j]
			j++
		}
	}
	if i < n {
		copy(nums3[i+j:], nums1[i:])
	}
	if j < m {
		copy(nums3[i+j:], nums2[j:])
	}

	if (n+m)%2 == 0 {
		median := float64(nums3[(n+m)/2]+nums3[(n+m)/2-1]) / 2.0
		return median
	} else {
		median := float64(nums3[(n+m)/2])
		return median
	}
}
