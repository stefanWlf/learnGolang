package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1, s2 := puppy.Bark(), puppy.Barks()

	fmt.Printf("%v \n %v \n", s1, s2)
	fmt.Println(puppy.BigBark(), puppy.BigBarks())
}
