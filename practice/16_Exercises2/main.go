package main

import (
	"fmt"
)

func main() {
	// for z := 0; z < 100; z++ {
	// 	x := rand.Intn(10)
	// 	y := rand.Intn(10)

	// 	fmt.Printf("The value of x is %v\nThe value of y is %v\n", x, y)
	// 	switch {
	// 	case x < 4 && y < 4:
	// 		fmt.Println("Both x and y are less than 4")
	// 	case x > 6 && y > 6:
	// 		fmt.Println("x and y are greater than 6")
	// 	case x >= 4 && x <= 6:
	// 		fmt.Println("x is between 4 and 6")
	// 	case y != 5:
	// 		fmt.Println("y isn't 5")
	// 	default:
	// 		fmt.Println("No specific criteria are met")
	// 	}
	// }

	//  for i := 0; i < 42; i++ {
	// 	x := rand.Intn(5)
	// 	switch x {
	// 	case 0:
	// 		fmt.Printf("Iteration is %v and value of x is %v\n", i, x)
	// 	case 1:
	// 		fmt.Printf("Iteration is %v and value of x is %v\n", i, x)
	// 	case 2:
	// 		fmt.Printf("Iteration is %v and value of x is %v\n", i, x)
	// 	case 3:
	// 		fmt.Printf("Iteration is %v and value of x is %v\n", i, x)
	// 	case 4:
	// 		fmt.Printf("Iteration is %v and value of x is %v\n", i, x)
	// 	}
	var i int = 10
	for i >= 0 {
		fmt.Printf("countdown %v\n", i)
		i--
	}

}
