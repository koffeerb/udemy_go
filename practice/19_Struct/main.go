package main

import "fmt"

type person struct {
	first    string
	last     string
	age      int
	icecream []string
}

type human struct {
	person
	isHuman bool
}

func main() {
	rw := human{
		person: person{
			first:    "Richie",
			last:     "Walsh",
			age:      38,
			icecream: []string{"Mint", "Pistachio"},
		},
		isHuman: true,
	}

	ak := human{
		person: person{
			first:    "Ali",
			last:     "Kavanagh",
			age:      37,
			icecream: []string{"Rum&Raisin", "Lemon Sorbet"},
		},
		isHuman: true,
	}
	//Anonymous struct not embedded struct
	tk := struct {
		first   string
		last    string
		age     int
		species string
	}{
		first:   "Taba",
		last:    "Kavanagh",
		age:     3,
		species: "cat",
	}

	// fmt.Println(rw)
	// fmt.Println(ak)
	fmt.Println(tk)

	// fmt.Printf("%#v\n", rw)
	// fmt.Printf("%#v\n", rw.person)
	// for _, v := range ak.icecream {
	// 	fmt.Println(ak.first, "'s favourite icecream is", v)
	// }
	fmt.Printf("%T\n", tk)

	m := map[string]human{
		rw.last: rw,
		ak.last: ak,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.icecream {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
