package main

import (
	"fmt"
)

func main() {

	// si := []string{"a", "b", "c"}
	// fmt.Println(si)

	// creates a slice and pre-builds a placeholder array up to capacity once it exceeds that it migrates it to a new slice of the same name
	xi := make([]int, 0, 10)

	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("------------")

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	fmt.Println("------------")

	xi = append(xi, 10, 11, 12, 13, 14, 15)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi)) // Increases capacity by multiple of initialised array
}
