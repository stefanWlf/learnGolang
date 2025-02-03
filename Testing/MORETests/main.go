package main

func Multiply(xi ...int) int {
	x := 1
	for _, v := range xi {
		x *= v
	}
	return x
}
