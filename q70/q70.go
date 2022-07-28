package q70

import "fmt"

func climbStairs(n int) int {
	s1, s2, result := 0, 0, 1
	for i := 1; i <= n; i++ {
		s1 = s2
		s2 = result
		result = s1 + s2
	}

	return result
}

func Test() {
	n := 3
	result := climbStairs(n)
	fmt.Println(result)
}
