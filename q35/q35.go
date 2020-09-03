package q35

import "fmt"

/**
https://leetcode-cn.com/problems/search-insert-position/
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
示例 1:
输入: [1,3,5,6], 5
输出: 2
*/

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	for i := 0; i < l; i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}

	return l
}

func Test() {
	nums := []int{1, 3, 5, 6}
	result := searchInsert(nums, 2)
	fmt.Println(result)
}
