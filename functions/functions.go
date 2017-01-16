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

func returnVariables(a, b string) (x, y int) {
	x = 5
	y = 3
	if a == "a" {
		x = y + 2
	}

	if b == "c " {
		y = y + x
	}

	return
}

func main() {
	fmt.Println(math.Pi)                               // Called of Identifier of package
	fmt.Println(add(2, 3))                             // Called functions with basic declaration of parameters
	fmt.Println(addOmmitType(4, 5))                    //Called functions with short declarations of parameters
	fmt.Println(multipleResults("Hello", "Stivenson")) // Called functions with multiple results

	x, y := multipleResults("Hello", "Stivenson") // Called functions for get two results in different variables
	fmt.Println(y, x)                             // Two result printing
	returnVariables("a", "b")                     // Called to function with return variables
}
