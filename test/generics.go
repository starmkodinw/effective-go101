package test

import "fmt"

func Generics[T any](data T) {
	fmt.Println("data : ", data)
}

func GenericsSlice[T any](data []T) {
	fmt.Println("len data : ", len(data))
}

func GenericsComparable() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15)) //2

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello")) //-1
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
