package q283

func moveZeroes(nums []int) {
	left := 0
	right := 0

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//func moveZeroes(nums []int) {
//	count := 0
//	for i := 0; i < len(nums); i++ {
//		if nums[i] != 0 {
//			nums[count] = nums[i]
//			count++
//		}
//	}
//
//	for i := count; i < len(nums); i++ {
//		nums[i] = 0
//	}
//}

func Test() {
	nums := []int{1, 2, 0, 0, 3}
	moveZeroes(nums)
}
