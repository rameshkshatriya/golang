package main

import "fmt"

func main() {
	// var day int
	// fmt.Println("enter a number: ")
	// fmt.Scanf(&day)

	switch day := 4; day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("invalid number")

	}

	var value int = 2
	switch {
	case value == 1:
		fmt.Println("moin ali khan")
	case value == 2:
		fmt.Println("meghanath prasad")
	case value == 3:
		fmt.Println("pannala madhusudhan reddy")
	}

	var letter string = "three"
	switch letter {
	case "one":
		fmt.Println("golang")
	case "two", "three":
		fmt.Println("python", "c")
	case "four":
		fmt.Println("java")

	}

}
