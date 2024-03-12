package test

import "fmt"

func Pointer() {
	type MyStruct struct {
		Value *int
	}

	x := 10
	y := 20
	a := MyStruct{Value: &x}
	b := a
	c := &a
	b.Value = &y
	d := b
	e := c
	f := &b
	*d.Value = *a.Value + *b.Value
	fmt.Println(*a.Value, *b.Value, *c.Value, *d.Value, *e.Value, *f.Value)
}

func Pointer2() {
	zz := 10
	fmt.Println(zz, &zz)

	var zzz *int = &zz
	fmt.Println(zzz, *zzz)

	new := zzz
	fmt.Println(new, *new)

	*new = 20
	fmt.Println(zz, *zzz, *new)
}

func Pointer3() {
	// การเข้าถึงและแก้ไขค่าผ่าน Pointer:

	// Address
	var num int = 10
	fmt.Println("Address of num:", &num) // 0xc0000a8008

	// Pointer
	var ptr *int = &num
	fmt.Println("Value of pointer:", *ptr) // 10

	// แก้ไขค่าผ่าน pointer
	*ptr = 20
	fmt.Println("New value of num:", num, *ptr) // 20
}

func Pointer4() {
	// เปรียบเทียบ Pointer:
	var num1 int = 10
	var num2 int = 10

	ptr1 := &num1
	ptr2 := &num2

	// เปรียบเทียบค่าของ pointer
	if ptr1 == ptr2 {
		s := fmt.Sprintf(`ptr1(%v) และ ptr2(%v) ชี้ไปที่ address เดียวกัน`, ptr1, ptr2)
		fmt.Println(s)
	} else {
		s := fmt.Sprintf(`ptr1(%v) และ ptr2(%v) ชี้ไปที่ address ต่างกัน`, ptr1, ptr2)
		fmt.Println(s)
	}

	// เปรียบเทียบค่าที่ pointer ชี้ไป
	if *ptr1 == *ptr2 {
		s := fmt.Sprintf(`ค่าที่ ptr1(%v) และ ptr2(%v) ชี้ไปนั้นเท่ากัน`, *ptr1, *ptr2)
		fmt.Println(s)
	} else {
		s := fmt.Sprintf(`ค่าที่ ptr1(%v) และ ptr2(%v) ชี้ไปนั้นต่างกัน`, *ptr1, *ptr2)
		fmt.Println(s)
	}
}

func Pointer5() {
	var ptr *int

	// ptr เป็น Nil pointer
	if ptr == nil {
		fmt.Println("ptr เป็น Nil pointer")
	}

	// กำหนดค่าให้ ptr
	num := 10
	ptr = &num
	fmt.Println(ptr, *ptr)
}

func Pointer6() {
	// Pointer to struct
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "John Doe", Age: 30}
	ptr := &person

	fmt.Println("ชื่อ:", (*ptr).Name, ptr.Name) // John Doe
	fmt.Println("อายุ:", ptr.Age)               // 30
}

func Pointer7() {
	// ประกาศ Pointer ไปยัง slice:
	var slice []int = []int{1, 2, 3, 4, 5}
	ptr := &slice

	// เข้าถึง element แรก
	fmt.Println("ค่า element:", (*ptr)[0])

	// เปลี่ยนค่า element แรก
	(*ptr)[0] = 10
	fmt.Println("ค่า element แรกหลังการเปลี่ยนแปลง:", (*ptr)[0]) // 10
}

func Pointer8() {
	// Dereferencing Pointer:
	var num int = 10
	ptr := &num

	// Dereferencing pointer เพื่อเข้าถึงค่า
	fmt.Println("ค่าของ num:", *ptr) // 10
}

func Pointer9() {
	var slice []int = []int{1, 2, 3, 4, 5}
	ptr := &slice

	// เพิ่ม element
	*ptr = append(*ptr, 6)

	// ลบ last element
	*ptr = (*ptr)[:len(*ptr)-1]

	fmt.Println("slice ใหม่:", *ptr) // [1 2 3 4 6]
}

func Pointer10() {
	var num1 int = 10
	var num2 int = 20

	swap(&num1, &num2)

	fmt.Println("num1:", num1) // 20
	fmt.Println("num2:", num2) // 10
}

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
