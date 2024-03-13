package test

import "fmt"

func Any() {
	var a any = 5
	a = "A"
	var b interface{} = 10
	b = "B"
	fmt.Println(a, b)
}
