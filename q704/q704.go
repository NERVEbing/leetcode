package q704

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		midValue := nums[mid]

		if target == midValue {
			return mid
		}
		if target > midValue {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func Test() {
	result := search([]int{-1, 0, 3, 5, 9, 12}, 12)
	fmt.Println(result)
}
