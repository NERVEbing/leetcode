package q21

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prevHead := &ListNode{Val: -1}
	prev := prevHead

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}

	return prevHead.Next
}

func Test() {
	// 1->2->4
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	// 1->3->4
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	result := mergeTwoLists(l1, l2)

	for {
		if result == nil {
			break
		}
		// 1->1->2->3->4->4
		fmt.Println(result.Val)
		result = result.Next
	}
}
