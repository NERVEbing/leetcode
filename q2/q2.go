package q2

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// i:是否超过10,n:两数之和
	var i, n int
	// head节点
	head := &ListNode{Val: 0}
	// 当前节点
	curr := head

	for {
		if i > 0 {
			n = l1.Val + l2.Val + 1
		} else {
			n = l1.Val + l2.Val
		}

		if n >= 10 {
			curr.Next = &ListNode{Val: n - 10}
			i = 1
		} else {
			curr.Next = &ListNode{Val: n}
			i = 0
		}

		curr = curr.Next

		if l1.Next == nil && l2.Next == nil {
			if i == 1 {
				curr.Next = &ListNode{Val: 1}
			}
			break
		}
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}

		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}

	return head.Next
}

func Test() {
	// [2,4,3]
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	// [5,6,4]
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := addTwoNumbers(l1, l2)

	for {
		if result == nil {
			break
		}
		// [7,0,8]
		fmt.Println(result.Val)
		result = result.Next
	}
}
