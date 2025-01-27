package main

import (
	"bytes"
	"fmt"
)

func main() {

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("World!")
	fmt.Println(b.String())
	b.Reset()
	fmt.Println(b.String())
	b.Write([]byte("Hello World! 2.0"))
	fmt.Println(b.String())
}
