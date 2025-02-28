package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	var l2 ListNode
	l1Ptr := &l1
	l2Ptr := &l2
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	rowStrs := strings.Fields(line)

	current1 := l1Ptr
	current2 := l2Ptr

	for i, numStr := range rowStrs {
		val, _ := strconv.Atoi(numStr)

		if i == 0 {
			current1.Val = val
		} else {
			current1.Next = &ListNode{Val: val}
			current1 = current1.Next
		}
	}
	scanner.Scan()
	line = scanner.Text()
	rowStrs = strings.Fields(line)
	for i, numStr := range rowStrs {
		val, _ := strconv.Atoi(numStr)

		if i == 0 {
			current2.Val = val
		} else {
			current2.Next = &ListNode{Val: val}
			current2 = current2.Next
		}
	}
	ptr := addTwoNumbers(l1Ptr, l2Ptr)
	for ptr != nil {
		fmt.Print(ptr.Val)
		if ptr.Next != nil {
			fmt.Print(" ")
		}
		ptr = ptr.Next
	}
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		val1 := 0
		if l1 != nil {
			val1 += l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 += l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}
