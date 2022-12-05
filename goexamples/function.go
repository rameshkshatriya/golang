package main

import "fmt"

func a(a, b int) int {
	sum := a + b
	return sum
}

func b(length, breadth int) int {
	mul := length * breadth
	return mul
}

func main() {
	fmt.Printf("adding the values a + b = %d\n", a(10, 20))

	fmt.Printf("multiplying the values a * b = %d", b(2, 3))
}
