package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	age := m["James"]
	fmt.Println(age)

	if test, ok := m["Moneypenny"]; ok {
		fmt.Println("There is one with that name, and he/she is", test, "years old")
	}

	if test, ok := m["DINKELBERG"]; !ok {
		fmt.Println("There is no one with that name. Her is the zero value", test)
	}

}
