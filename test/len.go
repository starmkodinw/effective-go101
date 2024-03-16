package test

import "fmt"

func Len() {
	s := []int{1, 2, 3, 4, 5, 6}

	s = s[:3]

	fmt.Println(len(s), cap(s), s)
}
