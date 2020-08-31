package q7

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	rn := 0
	for x != 0 {
		rn = rn*10 + x%10
		x /= 10
	}

	if rn > math.MaxInt32 || rn < math.MinInt32 {
		return 0
	}

	return rn
}

func Test() {
	result := reverse(1234)
	fmt.Println(result)
}
