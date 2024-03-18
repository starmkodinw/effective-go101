package test

import "fmt"

func Clear() {
	a := []int{1, 2, 3, 4, 5}
	clear(a)
	fmt.Println(a)

	mapS := map[string]string{
		"1": "one",
	}
	clear(mapS)
	fmt.Println(mapS)

	// Create a channel with buffer size 5
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println(len(ch)) // 5
	for len(ch) > 0 {    //clear buffer
		<-ch
	}
	fmt.Println(len(ch)) // 0
}
