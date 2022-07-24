package q121

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	max := 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}

func Test() {
	prices := []int{2, 1, 4}
	result := maxProfit(prices)
	fmt.Println(result)
}
