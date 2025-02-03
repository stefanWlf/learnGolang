package main

import (
	"fmt"

	quote "github.com/stefanWlf/learnGolang/TestingHandOnExercise/CountWords/Quote"
	word "github.com/stefanWlf/learnGolang/TestingHandOnExercise/CountWords/Words"
	dog "github.com/stefanWlf/learnGolang/TestingHandOnExercise/Dog"
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
