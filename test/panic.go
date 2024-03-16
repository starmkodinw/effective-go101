package test

import "fmt"

func Panic() {
	defer func() {
		if err := recover(); err != nil {
			switch {
			case err == 1:
				fmt.Println("Panic : ", err)
				fmt.Println("Send message to mark zuckerberg")
			default:
				fmt.Println("Panic default")
			}
		}
	}()

	panic(1)
	// panic(2)
}
