package q142

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast != nil && fast.Next != nil {
		for head != slow {
			head = head.Next
			slow = slow.Next
		}

		return head
	}

	return nil
}

func Test() {
	head := &ListNode{
		Val:  3,
		Next: nil,
	}

	pos := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  -4,
				Next: nil,
			},
		},
	}

	pos.Next.Next.Next = pos
	head.Next = pos

	fmt.Println(head)
	fmt.Println(head.Next)
	fmt.Println(head.Next.Next)
	fmt.Println(head.Next.Next.Next)
	fmt.Println(head.Next.Next.Next.Next)
	fmt.Println(head.Next.Next.Next.Next.Next)
	fmt.Println(head.Next.Next.Next.Next.Next.Next)
	fmt.Println("------")

	result := detectCycle(head)

	fmt.Println(result)
	//for {
	//	if result == nil {
	//		break
	//	}
	//	fmt.Println(result.Val)
	//	result = result.Next
	//}
}
