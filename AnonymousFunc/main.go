package main

import (
	"fmt"
	"math"
)

func returner() func() int {
	return func() int {
		return 42
	}
}

func add(a, b int) int {
	return a + b
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func powernator(a float64) func() float64 {
	var p float64
	return func() float64 {
		p++
		return math.Pow(a, p)
	}
}

func main() {
	func() {
		fmt.Println("anonymous func")
	}()

	a := func() {
		total := 0
		for i := 10; i > 0; i-- {
			total += i
		}
		fmt.Println(total)
	}
	a()

	f := returner()

	fmt.Println(f())

	fmt.Println(doMath(2, 2, add))

	pf := powernator(2)
	fmt.Println(pf())
	fmt.Println(pf())
	fmt.Println(pf())
	fmt.Println(pf())
	fmt.Println(pf())
	fmt.Println(pf())
	fmt.Println(pf())
	fmt.Println(pf())
}
