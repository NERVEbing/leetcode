package q35

import "fmt"

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high + 1
}

func Test() {
	nums := []int{-1, -2, 1, 3, 5, 6}
	result := searchInsert(nums, 0)
	fmt.Println(result)
}
