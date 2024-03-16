package test

import "fmt"

func New() {
	type S struct {
		a int
		b float64
	}
	s := new(S)
	(*s).a = 1
	fmt.Println(s, (*s).a, s.b)

	var s2 *S = &S{a: 1, b: 2.0}
	fmt.Println(s2, (*s2).a, (*s2).b)

	s3 := &S{
		a: 176,
	}
	fmt.Println(s3, (*s3).a, (*s3).b)
}
