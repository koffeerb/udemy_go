package main

import (
	"fmt"
	"sort"
)

/*
Slices use pointers to point at a specific address in memory for an underlying array

	Equating two slices mean they share the same pointers, so
	changing one value in the array changes will change the value for both slices at the same address
*/
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("------------------")
	a[2] = 100 // changing the value at array index 2 of slice a changes value for both slices
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("------ PART 2 ----------")

	c := []float64{0, 4, 2, 3, 1, 5, 6, 7}
	d := make([]float64, 8)

	copy(d, c) // Copies values to new array created using make

	fmt.Println(c)
	fmt.Println(d)

	// fmt.Println("------------------")
	// c[2] = 100 // this will only change the specified array
	// fmt.Println(c)
	// fmt.Println(d)

	fmt.Printf("\n------ PART 3 ----------\n")

	fmt.Println(medianOne(c))
	fmt.Println("after medianOne", c)

	fmt.Println(medianTwo(d))
	fmt.Println("after medianTwo", d)

}

func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	n := make([]float64, len(x))
	copy(n, x)
	// Copying allows for non-destructive analysis of underlying data.
	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
	}
	return (n[i-1] + n[i]) / 2

}
