package test

import "fmt"

func OrderEvaluation() {
	a := 1
	f := func() int { a++; return a }
	x := []int{a, f()} // f() ก่อน => x = [2, 2]

	y, ok := g(1, 2, f()), k() // ลำดับจะเป็น f(), g(1, 2, 3), k()
	fmt.Println(x, y, ok)

	// ลำดับการทำงาน
	// assignment, or return statement
	// all function calls
	// method calls
	// receive operations
	// and binary logical operations

	// Ex. y[f()], ok = g(z || h(), i()+x[j()], <-c), k()
	// f(), h() (if z evaluates to false), i(), j(), <-c, g(), and k()
}

func g(a, b, c int) int {
	return a + b + c
}

func k() bool {
	return true
}
