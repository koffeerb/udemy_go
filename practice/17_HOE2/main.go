package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println("before append", x)
	// x = append(x, 52)
	// fmt.Println("after append", x)
	// x = append(x, 53, 54, 55)
	// fmt.Println("after append", x)
	// y := []int{56, 57, 58, 59, 60}

	// x = append(x, y...)
	// fmt.Println("after append of second slice", x)
	x = append(x, x[3:6])

	fmt.Println(x)
}
