package functions_test

import (
	"fmt"
	"testing"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func TestClosures(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),    // 0 + 0, 1, 3
			neg(-2*i), // 0 + 0, -2,
		)
	}

	//sort.Sort()
}
