package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER OR EQUAL x which is %v \n", z, x)
		fmt.Printf("z is %v and that is LESS x which is %v \n", z, x)
	}
}
