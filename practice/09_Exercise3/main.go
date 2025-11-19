package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Printf("%v :Number is Odd\n", i)
		} else {
			fmt.Println("Number is even")
		}
	}
}
