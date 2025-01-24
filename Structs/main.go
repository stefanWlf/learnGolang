package main

import (
	"fmt"
)

func main() {
	usePerson()
	embedStruct()
	anonymousStruct()
}

type person struct {
	first string
	last  string
	favIC []string
}

func usePerson() {
	p1 := person{
		first: "Stefan",
		last:  "Wolf",
		favIC: []string{"Vanilla", "Chocolate"},
	}

	test := person{
		first: "Willy",
		last:  "Willbo",
		favIC: []string{"Fish"},
	}

	for i, v := range p1.favIC {
		fmt.Println(i, v)
	}
	fmt.Println(test)

	m1 := map[string]person{
		p1.last:   p1,
		test.last: test,
	}

	fmt.Println(m1)

	for k, v := range m1 {
		fmt.Printf("%v, %v \n", k, v.first)
		for a, b := range v.favIC {
			fmt.Println(a, "-", b)
		}
	}
}

type Engine struct {
	electric bool
}

type Vehicle struct {
	Engine
	published string
	model     string
	doors     int
	color     string
}

func embedStruct() {
	car := Vehicle{
		Engine:    Engine{electric: false},
		published: "01.02.2006",
		model:     "2006 Handmade, Custom-built, One-of-a-kind Race Car: lightning mcqueen",
		doors:     2,
		color:     "red with a yellow lightning strike",
	}
	car2 := Vehicle{
		Engine:    Engine{electric: false},
		published: "01.02.1957",
		model:     "1957 Haulital Hook'em",
		doors:     2,
		color:     "rusted brown",
	}

	fmt.Println(car)
	fmt.Println(car2)
}

func anonymousStruct() {
	person := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Peter",
		friends:   map[string]int{"Ulf": 53, "Q": 26},
		favDrinks: []string{"Tequilla", "Sex on the beach", "Guiness"},
	}

	fmt.Println(person)

	for k, v := range person.friends {
		fmt.Println(k, v)
	}
}
