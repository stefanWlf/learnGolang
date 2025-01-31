package main

func Subtract(xi ...int) int {
	x := 0
	for _, v := range xi {
		x += v
	}
	return x
}
