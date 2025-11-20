package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human struct {
	person
	isHuman bool
}

func main() {
	rw := human{
		person: person{
			first: "Richie",
			last:  "Walsh",
			age:   38,
		},
		isHuman: true,
	}

	ak := human{
		person: person{
			first: "Ali",
			last:  "Kavanagh",
			age:   37,
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

	fmt.Println(rw)
	fmt.Println(ak)
	fmt.Println(tk)

	fmt.Printf("%#v\n", rw)
	fmt.Printf("%#v\n", rw.person)
	fmt.Println(ak.isHuman)
	fmt.Printf("%T", tk)
}
