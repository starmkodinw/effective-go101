package test

import "fmt"

func AddressOperators() {
	var num *int // ประกาศ pointer `num`

	// `num` ไม่มี Address ของตัวแปร
	fmt.Println(*num) // panic: runtime error: invalid memory address or nil pointer dereference

	// nil pointer dereference
	num = nil
	fmt.Println(*num) // panic: runtime error: invalid memory address or nil pointer dereference

	// ใช้ safeDereference แทนการ dereference ตรงๆ
	value := safeDereference(num)
	fmt.Println("Value:", value) // Value: 0
}

func safeDereference(ptr *int) int {
	if ptr != nil {
		return *ptr
	}
	return 0
}
