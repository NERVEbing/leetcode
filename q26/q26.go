package q26

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[n] {
			n++
			nums[n] = nums[i]
		}
	}

	return n + 1
}

func Test() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums)
	fmt.Println(result)
	fmt.Println("-----")
	for i := 0; i < result; i++ {
		fmt.Println(nums[i])
	}
}
