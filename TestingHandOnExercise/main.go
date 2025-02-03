package main

import (
	"fmt"

	dog "github.com/stefanWlf/learnGolang/TestingHandOnExercise/Dog"
	quote "github.com/stefanWlf/learnGolang/TestingHandOnExercise/Quote"
	word "github.com/stefanWlf/learnGolang/TestingHandOnExercise/Word"
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
}
