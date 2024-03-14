package test

import "fmt"

func IntegerOverflow() {
	// type int16 เก็บค่า -32,768 ถึง 32,767
	var num int16 = 32767

	result, err := checkedAdd(num, 1)
	if err != nil {
		fmt.Println("Overflow:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func checkedAdd(x, y int16) (int16, error) {
	// ถ้าเกิด overflow => return error
	if x > 0 && y > 0 && x+y < 0 {
		return 0, fmt.Errorf("integer overflow")
	}
	if x < 0 && y < 0 && x+y >= 0 {
		return 0, fmt.Errorf("integer overflow")
	}
	return x + y, nil
}
