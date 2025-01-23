package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guiness", "Vanilla"}

	xxs := [][]string{jb, jm}

	fmt.Println(xxs)
}
