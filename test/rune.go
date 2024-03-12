package test

import "fmt"

func Rune() {
	var a = 'a'
	var b = rune('s')
	fmt.Println(rune(a), b, string(b))
}
