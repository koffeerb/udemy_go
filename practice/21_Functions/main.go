package main

import "fmt"

func main() {
	foo()
	bar("Richie")
	s := whizz("Ali")
	fmt.Println(s)
	s1, n := dogYears("Ali", 38)
	fmt.Println(s1, n)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := sum(xi...)
	fmt.Println(x)

}

func foo() {
	fmt.Println("From foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func whizz(s string) string {
	return fmt.Sprint("My name is ", s)

}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n

}
