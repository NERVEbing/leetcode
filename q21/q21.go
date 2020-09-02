package q21

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var result *ListNode

	if l1.Val > l2.Val {
		result = l2
		result.Next = mergeTwoLists(l1, l2.Next)
	} else {
		result = l1
		result.Next = mergeTwoLists(l1.Next, l2)
	}

	return result
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
