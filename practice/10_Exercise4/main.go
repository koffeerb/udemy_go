package main

import "fmt"

func main() {
	// xi := []int{42, 43, 44, 45, 46, 47}

	// for i, v := range xi {
	// 	fmt.Printf("index %v\t value %v\n", i, v)
	// }
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key is %v\t value is %v\n", k, v)
	}

	age := m["James"]
	fmt.Println("James is ", age)

	if _, ok := m["Q"]; !ok {
		fmt.Println("There is nobody by that name")
	}
}
