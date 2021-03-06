package main

import (
	"fmt"
)

// Function for square root with newton method
func Sqrt(x float64) float64 {
	z, n := float64(0), float64(1)
	for float32(z) <= float32(x) {
		n += 0.001
		z = float64(n * n)
	}
	return n
}

func main() {
	fmt.Println(Sqrt(2))
}
