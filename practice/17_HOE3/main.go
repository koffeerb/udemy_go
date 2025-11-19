package main

import "fmt"

func main() {
	usa := make([]string, 0, 50)
	usa = append(usa, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")
	fmt.Println(usa)
	fmt.Println(len(usa))
	fmt.Println(cap(usa))
	fmt.Println("-----------------------")

	for i := 0; i < len(usa); i++ {
		fmt.Println(i, usa[i])
	}

	fmt.Printf("\n------ PART 2 ----------\n")

	jb := []string{"James", "Bond", "Shaken not stirred"}
	jm := []string{"Miss", "MoneyPenney", "M is waiting"}
	spies := [][]string{jb, jm}

	for i, v := range spies {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}

	}

}
