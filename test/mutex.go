package test

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex // using Mutex

// using RWMutex
var rwMutex sync.RWMutex

func increment() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func readCounter() int {
	time.Sleep(100 * time.Millisecond)
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return counter
}
func Sec() {
	for i := 0; i < 10; i++ {
		go increment()
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Counter:", readCounter())
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Final counter:", counter)
}
