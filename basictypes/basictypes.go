package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) // imaginary part
	sqrt              = cmplx.Sqrt(2)        // is return and initialized
	f32    float32    = 0.04
	f64    float64    = 0.04
)

func main() {

	const f = "%T(%v) (%#v) \n" // All verbs: https://golang.org/pkg/fmt/
	fmt.Printf(f, ToBe, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt, MaxInt)
	fmt.Printf(f, z, z, z)
	fmt.Println(z)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(sqrt)
}
