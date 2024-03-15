package test

import "fmt"

func Conversions() {
	//  Implicit conversions
	var numb = 10 // num แปลงเป็น int โดยอัตโนมัติ
	fmt.Println(numb)

	// Explicit conversions โดยใช้ตัวดำเนินการ
	var floatNum float64 = 3.14
	var num int = int(floatNum) // แปลง float64 เป็น int
	fmt.Println(num)
}
