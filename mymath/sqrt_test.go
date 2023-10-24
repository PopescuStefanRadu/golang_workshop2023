package mymath_test

import (
	"esol/mymath"
	"fmt"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	for i := 1; i < 1500; i++ {
		t.Run(fmt.Sprintf("Testing for input: %v", i), func(t *testing.T) {
			want := math.Sqrt(float64(i))
			got := mymath.Sqrt(float64(i))
			if want != got {
				t.Fatalf("wanted: %v, got: %v", want, got)
			}
		})
	}
}
