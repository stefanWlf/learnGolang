package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123456"
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pw)

	err = bcrypt.CompareHashAndPassword(pw, []byte("12345678"))
	if err != nil {
		fmt.Println(err)
	}
}
