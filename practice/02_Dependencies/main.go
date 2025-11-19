package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Printf("%s", puppy.Bark()+"\n")
	fmt.Printf("%s", puppy.Barks()+"\n")
	fmt.Printf("%s", puppy.BigBark()+"\n")
}
