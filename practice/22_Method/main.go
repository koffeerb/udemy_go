package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func main() {

	p1 := person{
		first: "Richie",
		last:  "Walsh",
		age:   36,
	}

	p2 := person{
		first: "Ali",
		last:  "Kavanagh",
		age:   37,
	}

	p1.speak()
	p2.speak()
}
