package test

import (
	"fmt"
	"time"
)

func First() {
	// initialize
	ch := make(chan interface{})
	num := 0
	str := "*"

	// goroutine 1
	go func() {
		for {
			ch <- num
			num++
			time.Sleep(1 * time.Second)
		}
	}()

	// goroutine 2
	go func() {
		for {
			ch <- str
			str += "*"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// main goroutine
	for {
		fmt.Println("Value received:", <-ch)
	}
}
