package main

import (
	"fmt"
	"math/rand"
)

func init() {

	fmt.Println("This is the function initializing.")
}

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("The value of x is %v\nThe value of y is %v\n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("Both x and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is between 4 and 6")
	} else if y != 5 {
		fmt.Println("y isn't 5")
	} else {
		fmt.Println("No specific criteria are met")
	}

}
