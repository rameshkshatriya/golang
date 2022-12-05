package main

import "fmt"

func main() {
	a := 300

	if a == 200 {
		fmt.Printf("the value of a is 200")
	} else if a == 300 {
		fmt.Printf("the value of a is 300")
	} else if a == 400 {
		fmt.Printf("the value of a is 400")
	} else {
		fmt.Printf("not matching")
	}

}
