package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	y := rand.Intn(11)

	fmt.Printf("x: %v \t y: %v \n", x, y)

	switch {
	case x < 4 && y < 4:
		println("x and y are bot less then 4")
	case x > 6 && y > 6:
		println("x and y are both bigger than 6")
	case x >= 4 && x < 7:
		println("x is 4, 5 or 6")
	case y != 5:
		println("y is not 5")
	default:
		println("non of the above were triggered")
	}
}
