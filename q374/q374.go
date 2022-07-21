package q374

import (
	"fmt"
	"math/rand"
	"time"
)

func guessNumber(n int) int {
	low := 1
	high := n
	mid := low + (high-low)/2

	for guess(mid) != 0 {
		if guess(mid) == 1 {
			low = mid + 1
		}
		if guess(mid) == -1 {
			high = mid - 1
		}
		mid = low + (high-low)/2
	}

	return mid
}

func guess(n int) int {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(n) + 1
	if x > n {
		return 1
	}
	if x < n {
		return -1
	}

	return 0
}

func Test() {
	n := 10
	result := guessNumber(n)
	fmt.Println(result)
}
