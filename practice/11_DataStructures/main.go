// Arrays mostly used by GO internals - generally use slices instead.
package main

import (
	"fmt"
)

func main() {
	// var ni [7]int
	// ni[0] = 42
	// fmt.Printf("%#v \t\t %T\n", ni, ni)

	// ni2 := [4]int{55, 56, 57, 58}
	// fmt.Printf("%#v \t\t %T\n", ni2, ni2)

	// ns := [...]string{"Rooster", "Maris Piper", "Kerrs Pink"}
	// fmt.Printf("%#v \t %T\n", ns, ns)

	// fmt.Println(len(ni))
	// fmt.Println(len(ni2))
	// fmt.Println(len(ns))

	// var as = [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	// fmt.Println(as)
	// fmt.Println(len(as))

	food := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	// allow to use range and ignore the key value as a return
	for _, foodName := range food {
		fmt.Printf("%v\n", foodName)
	}
	fmt.Println("There are", len(food), "flavours of ice-cream")
	fmt.Println("----------------------------------------------")
	food = append(food, "Chocolate", "Mint", "Rum&Raisin")
	// You can also range over a slice by using
	// 	for _, foodName := range food[:4] {

	for _, foodName := range food {
		fmt.Printf("%v\n", foodName)
	}
	fmt.Println("There are", len(food), "flavours of ice-cream")
	fmt.Println("----------------------------------------------")
	// Slices of Slice
	fmt.Println(food[1:3]) // start at 2nd entry go to 4th
	fmt.Println(food[:2])  // start at 0th and go to 3rd entry
	fmt.Println(food[3:])  // Start at 4th and go to end
	fmt.Println(food[:])   // Return all
	fmt.Println("----------------------------------------------")

	// Deleting from Slice - append to set new values to slice, slice to position in array you want to remove, then unfurl from the next position to the end
	fmt.Println("There are", len(food), "flavours of ice-cream")
	fmt.Println("----------------------------------------------")

	food = append(food[:2], food[:3]...)
	for _, foodName := range food {
		fmt.Printf("%v\n", foodName)
	}
	fmt.Println("There are", len(food), "flavours of ice-cream")

}
