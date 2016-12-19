package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func addOmmitType(x, y int) int {
	return x + y
}

func multipleResults(x, y string) (string, string) {
	return x, y
}

func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(2, 3))
	fmt.Println(addOmmitType(4, 5))
	fmt.Println(multipleResults("Hello", "Stivenson"))

	x, y := multipleResults("Hello", "Stivenson")
	fmt.Println(y, x)

}
