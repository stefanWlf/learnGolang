package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person{
		First: "Uwe",
		Last:  "Müller",
		Age:   32,
	}
	p2 := person{
		First: "Jochen",
		Last:  "Mustermann",
		Age:   12,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(bs)

	var people2 []person

	bs = []byte(bs)
	err = json.Unmarshal(bs, &people2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n %v", people2)
}
