package test

import "fmt"

func TypeInference() {
	// Type inference หมายถึง กลไกที่ compiler สามารถอนุมานชนิดข้อมูล (type) ของตัวแปรได้โดยอัตโนมัติ

	num := 7                       //num: int
	platform := 9.75               //platform: float64
	message := "Expecto Patronum!" // message: string
	fmt.Println(num, platform, message)
}

func TypeUnification() {
	// Type unification หมายถึง กลไกที่ compiler สามารถแปลง ให้เป็น type เดียวกัน เพื่อใช้ใน operation หรือ assignment

	// Compiler แปลง `num` (int) เป็น float64 เพื่อใช้ใน operation `+`
	num := 1 + 2.5
	fmt.Println(num)
}
