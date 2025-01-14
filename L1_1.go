package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a Action) PerformAction() {
	fmt.Println("Performing an action...")
	a.Greet()
}

func main() {
	action := Action{
		Human: Human{
			Name: "Alice",
			Age:  30,
		},
	}

	action.PerformAction()
}
