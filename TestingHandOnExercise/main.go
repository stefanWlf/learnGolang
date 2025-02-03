package main

import (
	"fmt"

	word "github.com/stefanWlf/learnGolang/TestingHandOnExercise/CountWords"
	quote "github.com/stefanWlf/learnGolang/TestingHandOnExercise/CountWords/Quote"
	dog "github.com/stefanWlf/learnGolang/TestingHandOnExercise/Dog"
	mymath "github.com/stefanWlf/learnGolang/TestingHandOnExercise/myMath"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))

	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	xxi := gen()
	for _, v := range xxi {
		fmt.Println(mymath.CenteredAvg(v))
	}
}

func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
