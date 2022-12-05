package main

import "fmt"

func mul_div(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2
}

func main() {

	mul, _ := mul_div(2, 5)
	fmt.Println("2 * 5 = ", mul)
	// fmt.Println("2 / 4 = ", div)
}
