package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Number of cores:", runtime.NumCPU())
	fmt.Println("Name of the os:", runtime.GOOS)
	fmt.Println("Name of the architecture:", runtime.GOARCH)
	time.Sleep(time.Second * 10)
}
