package mymath

import "fmt"

var a = 1

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}
