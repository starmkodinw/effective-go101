package test

import "fmt"

func Defer() {
	defer func() {
		fmt.Println("defer 1")
		fmt.Println("defer 2")
	}()

	fmt.Println("main 1")
	fmt.Println("main 2")

	defer fmt.Println("defer 3")

	for i := 0; i <= 3; i++ {
		//defer 0
		//defer 1
		//defer 2
		//defer 3
		defer fmt.Println(i)
	}

	fmt.Println("f() : ", f())
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
