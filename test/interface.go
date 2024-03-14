package test

import "fmt"

// Define an interface
type MyInterface interface {
	Method1()
	Method2()
	EmbeddedInterface
}

type EmbeddedInterface interface {
	Method3()
	Method4()
}

// Implement the interface
type MyStruct struct {
	Name string
	Age  int32
}

// Method declaration ของ struct `MyStruct`
// method sets: ของ MyStruct มี Method1, Method2, Method3, Method4
// Method1 เป็น receiver pointer
// Method2, Method3, Method4 เป็น receiver value
func (s *MyStruct) Method1() {
	fmt.Println("Method1 called", s.Name)
}

func (s MyStruct) Method2() {
	fmt.Println("Method2 called", s.Age)
}

func (s MyStruct) Method3() {
	// implement Method3
}

func (s MyStruct) Method4() {
	// implement Method4
}

func NewMyStruct(name string) *MyStruct {
	return &MyStruct{
		Name: name,
		Age:  10,
	}
}

func TestInterface() {
	var i MyInterface = NewMyStruct("John") // MyStruct implement MyInterface
	i.Method1()
	i.Method2()
}

type PersonThai struct {
	Name string
}

func (p PersonThai) Hello(s string) {
	fmt.Println("Hello : ", p.Name, s)
}

func MethodExpression() {
	p := PersonThai{"John"}

	// Method expressions ใช้เรียก method ของ struct โดยไม่ต้องระบุชื่อ struct
	// เรียก method Hello() โดยใช้ method expression
	// สร้าง method value โดยใช้ method expression
	hello := p.Hello

	// เรียกใช้งาน method value
	hello("World")
}
