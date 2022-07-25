package q34

import (
	"fmt"
	"sort"
)

func searchRange(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}

	rightmost := sort.SearchInts(nums, target+1) - 1

	return []int{leftmost, rightmost}
}

func Test() {
	nums := []int{5, 7, 7, 9, 10}
	target := 7
	result := searchRange(nums, target)
	fmt.Println(result)
}
