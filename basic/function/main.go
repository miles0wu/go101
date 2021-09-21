package main

import (
	"fmt"
	"reflect"
)

var f1 = func() {
	fmt.Println("function 1")
}

func add(a int) func(b int) int {
	return func(b int) int { return a + b }
}

func main() {
	f1()

	value := func(input int) int {
		return input * 2
	}(5)

	fmt.Println(value)

	twoSumFunction := add(10)
	twoSum := add(10)(4)
	fmt.Println(reflect.TypeOf(twoSumFunction)) // output: function
	fmt.Println(twoSum)                         // output: 14
}
