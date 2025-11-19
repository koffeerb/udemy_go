package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("This will test statement-statement-idioms where a variable is set and an assertion is tested\n")
	x := 20
	fmt.Printf(" The Value of X is %v\n", x)

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v which is greater than or equal to x which is %v\n ", z, x)
	} else {
		fmt.Printf("z is %v which is less than x which is %v\n", z, x)
	}
}
