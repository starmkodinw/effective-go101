package test

import "fmt"

func ArithmeticOperators() {
	var num2 int = 4 // 100
	var num1 int = 2 // 010

	// Bitwise AND
	var result1 int = num1 & num2
	fmt.Println(result1) // 0

	// Bitwise OR
	var result2 int = num1 | num2
	fmt.Println(result2) // 6

	// Bitwise XOR => ต่างกัน 1
	var result3 int = num1 ^ num2
	fmt.Println(result3) // 6
}
