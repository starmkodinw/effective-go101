package test

import "fmt"

func Printing() {
	type T2 struct {
		a int
		b float64
		c string
	}
	t := &T2{7, -2.35, "abc\tdef"}
	p := "%"
	fmt.Printf("%vv ========> %v\n", p, t)   //value
	fmt.Printf("%v+v ========> %+v\n", p, t) //fields + value
	fmt.Printf("%v#v ========> %#v\n", p, t) //full Go syntax
	fmt.Printf("%vT ========> %T\n", p, t)   //type
}
