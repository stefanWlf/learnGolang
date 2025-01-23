package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Printf("i = %v \n", i)
	// }

	// for i := 0; i < 42; i++ {
	// 	x := rand.Intn(5)

	// 	if x%2 == 0 {
	// 		fmt.Printf("Iteration: %v \t x: %v, is even \n", i, x)
	// 	} else {
	// 		fmt.Printf("Iteration: %v \t x: %v, is odd \n", i, x)
	// 	}
	// }

	y := 0

	for y < 10 {
		fmt.Printf("Iteration: %v \n", y)
		y++
	}
}
