package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 *ListNode
	var l2 *ListNode
	l1 = &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergedList := mergeTwoLists(l1, l2)
	for mergedList != nil {
		fmt.Print(mergedList.Val, " ")
		mergedList = mergedList.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	op := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			op.Next = list1
			list1 = list1.Next
		} else {
			op.Next = list2
			list2 = list2.Next
		}
		op = op.Next
	}
	if list1 != nil {
		op.Next = list1
	}
	if list2 != nil {
		op.Next = list2
	}
	return dummy.Next
}

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists(l1.Next, l2)
//		return l1
//	}
//	l2.Next = mergeTwoLists(l1, l2.Next)
//	return l2
//}
