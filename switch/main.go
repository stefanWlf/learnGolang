package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {

		x := rand.Intn(251)
		
		switch {
		case x == 0:
			fmt.Printf("your value is %v that means: x is 0 \n", x)
		case x < 100 && x != 0:
			fmt.Printf("your value is %v that means: x is less than 100 \n", x)
		case x > 100 && x < 200:
			fmt.Printf("your value is %v that means: x is between 100 and 200 \n", x)
		case x > 200 && x != 250:
			fmt.Printf("your value is %v that means: x is greater than 200 and less then 250 \n", x)
		case x == 250:
			fmt.Printf("your value is %v that means: x is the maximum: 250! \n", x)
		default:
			fmt.Printf("your value is %v. That was not expected! \n", x)	
		}
	}
}
