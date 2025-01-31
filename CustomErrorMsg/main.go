package main

import "fmt"

type customErr struct {
	error string
	value int
}

func (err customErr) Error() string {
	return fmt.Sprintf("Folgender Fehler ist aufgetaucht: %v, Fehlercode: %d", err.error, err.value)
}

func main() {
	err := customErr{
		error: "ahhh lol error",
		value: 1,
	}
	foo(err)
}

func foo(err error) {
	fmt.Println("foo ist gelaufen", err)
}
