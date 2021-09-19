package main

import (
	"fmt"
	"reflect"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type Human struct {
	Name   string
	Gender Gender
}

type Car struct {
	Brand string `json:"brand"`
}

func main() {
	h1 := new(Human)
	h1.Name = "Miles"
	h1.Gender = GenderMale
	fmt.Println(h1)

	h2 := Human{Name: "Ann", Gender: GenderFemale}
	fmt.Println(h2)

	car1 := Car{Brand: "Benz"}
	car1Type := reflect.TypeOf(car1)
	fmt.Println(car1Type)
	brand := car1Type.Field(0)
	fmt.Println(brand)
	brandTag := brand.Tag.Get("json")
	fmt.Println(brandTag)
}
