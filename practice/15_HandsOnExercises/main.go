package main

import "fmt"

func main() {
	counting := [5]int{}
	// for i := range 5 {
	for i := 0; i < 5; i++ {
		counting[i] = i
	}
	fmt.Printf("%v\n%T\n", counting, counting)

	for _, value := range counting {
		fmt.Printf("The value is %v, The type is %T\n", value, value)
	}

	fmt.Printf("\n------ PART 2 ----------\n")

	midas := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, value := range midas {
		fmt.Printf("The value is %v, The type is %T\n", value, value)
	}

	fmt.Printf("\n------ PART 3 ----------\n")

	gold1 := midas[:5]
	gold2 := midas[5:]
	gold3 := midas[2:7]
	gold4 := midas[1:6]
	fmt.Println(gold1)
	fmt.Println(gold2)
	fmt.Println(gold3)
	fmt.Println(gold4)

}
