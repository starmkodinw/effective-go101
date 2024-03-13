package test

import "fmt"

const (
	// iota เริ่มต้นใหม่ใน const block นี้
	sunday = iota
	monday
	tuesday
)

const (
	// iota เริ่มต้นใหม่ใน const block นี้
	_ = iota // _ ละเว้นค่า
	north
	south
	east
	west
)

func Iota() {
	fmt.Println(sunday)  // 0
	fmt.Println(monday)  // 1
	fmt.Println(tuesday) // 2

	fmt.Println(north) // 1
	fmt.Println(south) // 2
	fmt.Println(east)  // 3
	fmt.Println(west)  // 4
}
