package main

import (
	"fmt"
	"reflect"
)

type A struct {
	PublicItem  string
	privateItem string
}

func (a A) Implode() string {
	return a.PublicItem + "-" + a.privateItem
}

func main() {
	myMap := make(map[string]string, 10)

	myMap["a"] = "b"

	t := reflect.TypeOf(myMap)
	fmt.Println("type is:", t)

	v := reflect.ValueOf(myMap)
	fmt.Println("value is:", v)

	a := A{"One", "2"}
	va := reflect.ValueOf(a)
	fmt.Println(va)

	for i := 0; i < va.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, va.Field(i))
	}

	for i := 0; i < va.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, va.Method(i))
	}

	fmt.Println(va.Method(0).Call(nil))
}
