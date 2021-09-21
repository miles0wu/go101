package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	FirstName string
	LastName  string
}

type Car struct {
	Factory string
	Model   string
}

func (h *Human) getName() string {
	return fmt.Sprintf("%s, %s", h.FirstName, h.LastName)
}

func (c *Car) getName() string {
	return c.Factory + "-" + c.Model
}

func main() {
	objects := []IF{
		&Human{FirstName: "Miles", LastName: "Wu"},
		&Car{Factory: "Toyota", Model: "Altis"},
	}

	for _, object := range objects {
		fmt.Println(object.getName())
	}
}
