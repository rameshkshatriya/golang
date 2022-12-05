package main

import "fmt"

func main() {
	var r int
	fmt.Print("enter a number: ")
	fmt.Scan(&r)
	for i := 0; i < r; i++ {
		for j := 0; j < r-i-1; j++ {
			fmt.Print(" ", " ")
		}
		for k := 0; k < 1*i+1; k++ {
			fmt.Print("*", " ")
		}
		fmt.Println()
	}
	for l := 0; l < r; l++ {
		for m := 0; m < r; m++ {
			fmt.Print(" ", " ")
		}
		for n := 0; n < 1*(r-l)-1; n++ {
			fmt.Print("*", " ")
		}
		fmt.Println()
	}
}
