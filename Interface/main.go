package main

import (
	"fmt"
	"math"
)

type square struct {
	length, width float64
}

type circle struct {
	radius float64
}

type shape interface {
	areaCalculation() float64
}

func (s square) areaCalculation() float64 {
	return s.width * s.length
}

func (c circle) areaCalculation() float64 {
	return c.radius * math.Pow(math.Pi, 2)
}
func Info(s shape) {
	fmt.Println("Your area is:", s.areaCalculation())
}

func main() {

	s1 := square{
		width:  25,
		length: 10,
	}
	c1 := circle{
		radius: 15,
	}

	Info(s1)
	Info(c1)

}
