package test

import "fmt"

func MinMax() {
	x := 10
	fmt.Println(min(x, 2))
	fmt.Println(max(float64(x), 1, 2.23))

	s := "A"
	fmt.Println(min(s, "B"))
	fmt.Println(max("", "foo", "bar"))
}
