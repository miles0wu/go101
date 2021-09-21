package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	person := Person{Name: "aaa"}
	fmt.Println(person)

	changeName(&person, "bbb")
	fmt.Println(person)

	cannotChangeName(person, "ccc")
	fmt.Println(person)
}

func changeName(person *Person, name string) {
	person.Name = name
}

func cannotChangeName(person Person, name string) {
	person.Name = name
}
