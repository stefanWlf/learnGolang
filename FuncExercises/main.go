package main

import (
	"fmt"
)

type user struct {
	first string
	age   int
}

func main() {
	fmt.Println("Anfang des Main blocks")
	fmt.Println("------------------------------")
	defer fmt.Println("-------------------------------------------")
	defer deferTest()

	x := foo(2, 40)
	y, z := bar(42, "Hello World")

	fmt.Println(x)
	fmt.Printf("int: %v, String: %v \n \n", y, z)

	g := goo(1, 2, 3)
	gg := gar([]int{1, 23, 5, 6, 89, 9})

	fmt.Println(g, gg)

	u1 := user{
		first: "Max",
		age:   42,
	}
	u1.Speak()

	fmt.Println("-------------------------------------------")


	fmt.Println("-------------------------------------------")
	fmt.Println("Ende des Main Blocks")
}

func deferTest() {
	fmt.Println("Jetzt läuft defer, da alles in der main func durchgelaufen ist")
}

func (u user) Speak() {
	fmt.Printf("\n Ich bin %v und %v Jahre alt \n", u.first, u.age)
}

func foo(a, b int) int {
	return a + b
}

func bar(a int, b string) (int, string) {
	return a, b
}

func goo(a ...int) int {
	t := 0
	for _, v := range a {
		t += v
	}
	return t
}

func gar(xa []int) int {
	x := 0
	for _, v := range xa {
		x += v
	}
	return x
}
