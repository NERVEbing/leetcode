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
	if nums == nil || len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
		if i+1 == len(nums) && nums[i] < target {
			return i + 1
		}
	}

	return 0
}

func Test() {
	nums := []int{1, 3, 5, 6}
	result := searchInsert(nums, 2)
	fmt.Println(result)
}
