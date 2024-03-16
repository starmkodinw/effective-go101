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
}
