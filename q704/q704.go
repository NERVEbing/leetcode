package q704

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high--
		} else {
			low++
		}
	}

	return -1
}

func Test() {
	result := search([]int{-1, 0, 3, 5, 9, 12}, 3)
	fmt.Println(result)
}
