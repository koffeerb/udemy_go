package main

import (
	"fmt"
	"math/rand"
)

func main() {
	potato := rand.Intn(15)
	switch {
	case potato == 10:
		fmt.Printf("potato is 10")
	case potato < 10:
		fmt.Printf("potato is less than 10")
	case potato > 10:
		fmt.Printf("potato is more than 10")
	default:
		fmt.Printf("Default for potato")
	}
	fmt.Printf("\nPotato is %v", potato)
}
