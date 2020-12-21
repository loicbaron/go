package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %v, value = %v\n", i, z)
	}
	return z
}

func main() {
	x := float64(4)
	fmt.Println(math.Sqrt(x))
	fmt.Println(sqrt(x))
}
