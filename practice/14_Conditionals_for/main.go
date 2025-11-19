package main

import (
	"fmt"
)

func main() {
	for x := 0; x <= 20; x++ {
		fmt.Printf("Counting to 20 \n")
		z := x / 2
		fmt.Printf("x is divisible by 2 %v times with remainder %v \n", z, x%2)
		for y := 0; y <= 3; y++ {
			fmt.Printf("x is %v and y is value %v \n", x, y)
		}
	}
}
