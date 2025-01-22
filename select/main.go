package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {

		x := rand.Intn(251)
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)
		ch4 := make(chan int)
		ch5 := make(chan int)

		go func() {
			if x != 0 {
				return
			}
			ch1 <- x
		}()

		go func() {
			if x > 100 && x == 0 {
				return
			}
			ch2 <- x
		}()

		go func() {
			if x < 100 && x > 200 {
				return
			}
			ch3 <- x
		}()

		go func() {
			if x < 200 && x == 250 {
				return
			}
			ch4 <- x
		}()
		go func() {
			if x != 250 {
				return
			}
			ch5 <- x
		}()

		select {
		case v1 := <-ch1:
			fmt.Printf("your value is %v that means: x is 0 \n", v1)
		case v2 := <-ch2:
			fmt.Printf("your value is %v that means: x is less than 100 \n", v2)
		case v3 := <-ch3:
			fmt.Printf("your value is %v that means: x is between 100 and 200 \n", v3)
		case v4 := <-ch4:
			fmt.Printf("your value is %v that means: x is greater than 200 and less then 250 \n", v4)
		case v5 := <-ch5:
			fmt.Printf("your value is %v that means: x is the maximum: 250! \n", v5)
		default:
			fmt.Printf("your value is %v. That was not expected! \n", x)
		}
	}
}
