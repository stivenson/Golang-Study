package main

import (
	"fmt"
)

func runIf(cat, dog string) string {
	if cat2 := dog; cat == "chispa" && cat2 != "Misifú" {
		return "She's chispa"
	}
	return "She's doesn't chispa"
}

func main() {
	fmt.Println(runIf("chispa", "Misifú"))
	fmt.Println(runIf("chispa", "Rasputín"))
}
