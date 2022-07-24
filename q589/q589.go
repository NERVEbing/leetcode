package q589

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var result []int
	var f func(*Node)

	f = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, c := range node.Children {
			f(c)
		}
	}
	f(root)

	return result
}

func Test() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}

	result := preorder(root)
	fmt.Println(result)
}
