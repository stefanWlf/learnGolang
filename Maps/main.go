package main

import "fmt"

func main() {
	m1 := createMap()
	m1 = appendMap(m1)
	deleteMap(m1)
}

func createMap() map[string]([]string) {
	names := map[string]([]string){
		`bond_james`:       []string{`Shaken, not stirred`, `martinis`, `fast cars`},
		`moneypenny_jenny`: []string{`intelligence`, `literature`, `computer science`},
		`no_dr`:            []string{`cats`, `ice cream`, `sunsets`},
	}

	return names
}

func appendMap(m map[string][]string) map[string][]string {
	m1 := m
	m1[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	/*
		for k, v := range m1 {
			fmt.Printf("%v \n\n", k)
			for i, v2 := range v {
				fmt.Printf("%v - %v \t \t %T\n\n", i, v2, v2)
			}
		}
	*/

	return m1
}

func deleteMap(m map[string][]string) {
	delete(m, "no_dr")
	for k, v := range m {
		fmt.Printf("%v \n\n", k)
		for i, v2 := range v {
			fmt.Printf("%v - %v \t \t %T\n\n", i, v2, v2)
		}
	}
}
