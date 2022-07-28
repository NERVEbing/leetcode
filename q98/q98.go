package q98

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var f func(*TreeNode, int, int) bool

	f = func(node *TreeNode, low int, high int) bool {
		if node == nil {
			return true
		}
		if node.Val <= low || node.Val >= high {
			return false
		}

		return f(node.Left, low, node.Val) && f(node.Right, node.Val, high)
	}

	return f(root, math.MinInt, math.MaxInt)
}

func Test() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := isValidBST(root)
	fmt.Println(result)
}
