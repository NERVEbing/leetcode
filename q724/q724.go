package q724

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	leftSum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if leftSum*2+nums[i] == sum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

func Test() {
	nums := []int{1, 7, 3, 6, 5, 6}
	result := pivotIndex(nums)
	fmt.Println(result)
}
