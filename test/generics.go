package test

import "fmt"

func Generics[T any](data T) {
	fmt.Println("data : ", data)
}

func GenericsSlice[T any](data []T) {
	fmt.Println("len data : ", len(data))
}
