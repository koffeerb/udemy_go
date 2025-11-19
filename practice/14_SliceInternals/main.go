package main

import "fmt"

/* Slices use pointers to point at a specific address in memory for an underlying array
   Equating two slices mean they share the same pointers, so
   changing one value in the array changes will change the value for both slices at the same address  */
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

	c := []int{0, 1, 2, 3, 4, 5}
	d := make([]int, 6)

	copy(d, c) // Copies values to new array created using make

	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("------------------")
	c[2] = 100 // this will only change the specified array
	fmt.Println(c)
	fmt.Println(d)

}
