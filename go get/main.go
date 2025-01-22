package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1, s2 := puppy.Bark(), puppy.Barks()
	
	fmt.Printf("%v \n %v", s1, s2)
}