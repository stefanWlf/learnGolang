package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello I'm", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	wg.Add(102)

	p1 := &person{
		"Uwe",
	}

	saySomething(p1)
	seconds := rand.Intn(1)
	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		fmt.Println("#1 I needed", (time.Second*time.Duration(seconds)) + 1, "seconds")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		fmt.Println("#2 I needed", (time.Second*time.Duration(seconds)) + 1, "seconds")
		wg.Done()
	}()
	var counter int

	// intentionally created a race condition
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
}
