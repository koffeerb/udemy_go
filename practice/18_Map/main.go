package main

import "fmt"

func main() {
	rt := map[string]int{
		"Ali":    38,
		"Richie": 37,
		"Taba":   4,
	}
	fmt.Println(rt)
	fmt.Printf("\nTaba is %v years old\n", rt["Taba"])
	fmt.Println("==================================")

	for k, v := range rt {
		println(k, v)
	}
	bw := make(map[string]int)
	bw["Gerry"] = 71
	bw["Mo"] = 68
	bw["Brendan"] = 33

	fmt.Println(bw)
	for k, v := range bw {
		println(k, v)
	}
}
