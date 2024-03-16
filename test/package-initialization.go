package test

import "fmt"

var (
	a9 = c9 + b9 // == 9
	b9 = f2()    // == 4
	c9 = f2()    // == 5
	d9 = 3       // == 5 after initialization has finished
)

func f2() int {
	d9++
	return d9
}

func PackageInitialization() {
	// the initialization order is d9, b9, c9, a9
	fmt.Println(a9, b9, c9, d9) // 9 4 5 5
}
