package main

import (
	"fmt"
)

func main() {

	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Miss", "Moneypenny", "Guinness", "Mint"}
	// fmt.Println(jb)
	// fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)

	fmt.Println("------------")
	// Cannot access individual indices of constituent slices.
	fmt.Println(len(xxs))
}
