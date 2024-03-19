package test

import (
	"fmt"
	"math"
)

func TestGithub() {
	// CTRL + SHIFT + I
	// /explain
	// /simplify
	// /fix
	// /tests

	fmt.Println("Before : ", "Hello")
	fmt.Println("After : ", fGithub("Hello"))

	count := 0
	result := math.Pow(2, 5) - 1
	hannoi(5, "A", "C", "B", &count)
	fmt.Println(count, "calls", result)
}

/*
Hello
World
*/
func fGithub(v string) string {
	return "World"
}

// solve the tower of hannoi problem
// n: number of disks
// from: source rod
// to: destination rod
// aux: auxiliary rod
func hannoi(n int, from, to, aux string, count *int) {
	*count++
	if n == 1 {
		fmt.Println("Move disk 1 from", from, "to", to)
		return
	}
	hannoi(n-1, from, aux, to, count)
	fmt.Println("Move disk", n, "from", from, "to", to)
	hannoi(n-1, aux, to, from, count)
}
