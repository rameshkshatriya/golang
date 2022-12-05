package main

import "fmt"

func main() {

	var a int = 24
	var b int = 32

	if a != b && a <= b {
		fmt.Println("True")
	} else {
		fmt.Println("false")
	}

	if a != b || a >= b {
		fmt.Println("True")
	} else {
		fmt.Println("false")
	}

	if !(a == b) {
		fmt.Println("True")
	}
}
