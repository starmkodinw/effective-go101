package test

import (
	"fmt"
)

func TestGithub() {
	// CTRL + SHIFT + I
	// /explain
	// /simplify
	// /fix
	// /tests

	fmt.Println("Before : ", "Hello")
	fmt.Println("After : ", fGithub("Hello"))
}

func fGithub(v string) string {
	return "World"
}
