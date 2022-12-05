package main

import (
	"fmt"
)

func main() {

	const a = "GFG"
	var b = "Greekforgreek"
	const c = "24"
	var d = "26"
	const e = 24
	var f = 26

	var concat = a + " " + b
	fmt.Println(concat)

	con := c + " " + d
	fmt.Println(con)

	add := e + f
	fmt.Println(add)

	fmt.Println(a == "GFG")
	fmt.Println(b < a)
}
