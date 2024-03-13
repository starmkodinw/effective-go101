package test

import "fmt"

func CompositeLiteral() {
	// Composite literal สำหรับ struct
	person := struct {
		Name string
		Age  int
	}{
		Name: "John Doe",
		Age:  30,
	}

	// Composite literal สำหรับ array
	numbers := [...]int{1, 2, 3, 4, 5}

	// Composite literal สำหรับ slice
	fruits := []string{"apple", "banana", "orange"}

	// Composite literal สำหรับ map
	ages := map[string]int{
		"John": 30,
		"Jane": 25,
	}

	// Function literal
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println(person, numbers, fruits, ages, sum(1, 2))
}
