package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This starts the init function")
}

func main() {
	for i := 0; i < 10; i++ {

		/*
		 Wrong use of the select statement, but I just wanted to use and try the select statement :)
		*/

		x := rand.Intn(251)

		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)
		ch4 := make(chan int)
		ch5 := make(chan int)

		go func() {
			if x == 0 {
				ch1 <- x
			}
		}()

		go func() {
			if x < 100 && x != 0 {
				ch2 <- x
			}
		}()

		go func() {
			if x > 99 && x < 200 {
				ch3 <- x
			}
		}()

		go func() {
			if x > 199 && x != 250 {
				ch4 <- x
			}
		}()
		go func() {
			if x == 250 {
				ch5 <- x
			}
		}()

		select {
		case v1 := <-ch1:
			fmt.Printf("your value is %v that means: x is the smallest possible number! \n", v1)
		case v2 := <-ch2:
			fmt.Printf("your value is %v that means: x is less than 100 \n", v2)
		case v3 := <-ch3:
			fmt.Printf("your value is %v that means: x is between 100 and 200 \n", v3)
		case v4 := <-ch4:
			fmt.Printf("your value is %v that means: x is greater than 200 and less then 250 \n", v4)
		case v5 := <-ch5:
			fmt.Printf("your value is %v that means: x is the maximum: 250! \n", v5)
		}
	}
}
