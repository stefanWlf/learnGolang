package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	get(c)
	send(c)
	fmt.Println("----------")
	c2 := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c2 <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c2)
	}
}

func get(c chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
}

func send(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
