package main

import (
	"fmt"
)

const Pi = 3.1416

// Numeric constans

const (
	Big   = 1 << 100  // Bits duplicate to 100
	Small = Big >> 99 // Bits reduce to 99
)

func needFloat(f float64) float64 {
	return f * 0.1
}

func main() {
	const cn1 = "a text"
	const cn2 = 2
	const cn3 = true
	fmt.Println(cn1, cn2, cn3, Pi)
	fmt.Println(needFloat(Big), Small)
}
