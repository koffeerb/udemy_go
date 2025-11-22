package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	mi5 bool
}

func (spy secretAgent) speak() {
	fmt.Println("I am a secret agent", spy.first)
}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}

// Anything with this method is also of type human
// Allows for polymorphism
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	spy1 := secretAgent{
		person: person{
			first: "Richie",
			last:  "Walsh",
			age:   36,
		},
		mi5: true,
	}

	p2 := person{
		first: "Ali",
		last:  "Kavanagh",
		age:   37,
	}

	// spy1.speak()
	// p2.speak()

	saySomething(spy1)
	saySomething(p2)
}
