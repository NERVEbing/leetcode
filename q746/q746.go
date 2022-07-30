package q746

import "fmt"

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	f := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 2; i <= n; i++ {
		dp[i] = f(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

func Test() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	result := minCostClimbingStairs(cost)
	fmt.Println(result)
}
