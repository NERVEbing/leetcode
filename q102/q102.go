package q102

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	arr := []*TreeNode{root}
	for i := 0; len(arr) > 0; i++ {
		result = append(result, []int{})
		var arrLevel []*TreeNode
		for n := 0; n < len(arr); n++ {
			node := arr[n]
			result[i] = append(result[i], node.Val)
			if node.Left != nil {
				arrLevel = append(arrLevel, node.Left)
			}
			if node.Right != nil {
				arrLevel = append(arrLevel, node.Right)
			}
		}
		arr = arrLevel
	}

	return result
}

func Test() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	result := levelOrder(root)
	fmt.Println(result)
}
