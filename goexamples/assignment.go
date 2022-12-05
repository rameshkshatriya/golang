package main

import "fmt"

func main() {

	a := 30
	b := 16

	a = b
	fmt.Println(a)

	a += b
	fmt.Println(a)

	a -= b
	fmt.Println(a)

	a *= b
	fmt.Println(a)

	a /= b
	fmt.Println(a)

	a %= b
	fmt.Println(a)
}
