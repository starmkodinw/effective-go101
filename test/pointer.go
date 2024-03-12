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
	// Address
	var num int = 10
	fmt.Println("Address of num:", &num) // 0xc0000a8008

	// Pointer
	var ptr *int = &num
	fmt.Println("Value of pointer:", *ptr, ptr, &num) // 10

	// แก้ไขค่าผ่าน pointer
	*ptr = 20
	fmt.Println("New value of num:", num, *ptr) // 20
}
