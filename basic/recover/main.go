package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println("message from defer:", err)
		}
	}()

	panic("panic triggered!!")
}
