// Arrays mostly used by GO internals - generally use slices instead.
package main

import (
	"fmt"
)

func main() {
	var ni [7]int
	ni[0] = 42
	fmt.Printf("%#v \t\t %T\n", ni, ni)

	ni2 := [4]int{55, 56, 57, 58}
	fmt.Printf("%#v \t\t %T\n", ni2, ni2)

	ns := [...]string{"Rooster", "Maris Piper", "Kerrs Pink"}
	fmt.Printf("%#v \t %T\n", ns, ns)

	fmt.Println(len(ni))
	fmt.Println(len(ni2))
	fmt.Println(len(ns))
}
