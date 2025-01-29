package main

import (
	"fmt"
	"math/rand"
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
	seconds := rand.Intn(5)
	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		fmt.Println("#1 I needed", (time.Second * time.Duration(seconds)), "seconds")
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		fmt.Println("#2 I needed", (time.Second * time.Duration(seconds)), "seconds")
		wg.Done()
	}()
	/*
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

		// fix with the atomic package
		var counter2 int64
		for i := 0; i < 100; i++ {
			go func() {
				atomic.AddInt64(&counter2, 1)
				fmt.Println(counter2)
				wg.Done()
			}()
		}
	*/
	// fix with mutex 
	var mu sync.Mutex
	var counter int
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			mu.Unlock()
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
}
