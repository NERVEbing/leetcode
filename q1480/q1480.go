package q1480

import "fmt"

func runningSum(nums []int) []int {
	last := nums[0]

	for i := 0; i < len(nums); i++ {
		if i > 0 {
			nums[i] = last + nums[i]
			last = nums[i]
		}
	}

	return nums
}

func Test() {
	nums := []int{3, 1, 2, 10, 1}
	result := runningSum(nums)
	fmt.Println(result)
}
