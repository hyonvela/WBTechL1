package main

import "fmt"

type Greet interface {
	hello(name string)
}

type Person struct{}

func (p Person) speak(text string) {
	fmt.Println(text)
}

type FriendlyPerson struct {
	person Person
}

func (p FriendlyPerson) hello(name string) {
	p.person.speak(fmt.Sprintln("Hello,", name))
}

func check(f Greet) {
	f.hello("Jhon")
	fmt.Println("Ok!")
}

func main() {
	person := Person{}

	adapter := FriendlyPerson{person}

	check(adapter)
}
