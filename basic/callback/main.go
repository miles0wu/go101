package main

import "fmt"

func main() {
	DoOperation(10, decrease)
}

func DoOperation(input int, callback func(int, int)) {
	callback(input, 1)
}

func decrease(a, b int) {
	fmt.Println(a - b)
}
