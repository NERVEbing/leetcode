package q27

import "fmt"

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}

	return n
}

func Test() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	result := removeElement(nums, 2)
	fmt.Println(result)
	fmt.Println("-----")
	for i := 0; i < result; i++ {
		fmt.Println(nums[i])
	}
}
