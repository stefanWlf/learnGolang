package main

import "fmt"

var x int = 42

const (
	_ = iota
	a = (10 * iota) + 1
	b
	c
)

func main() {
	y := 24
	fmt.Println(x * a - y / c)
}
