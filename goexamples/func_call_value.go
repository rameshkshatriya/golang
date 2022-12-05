package main

import "fmt"

func swap(x, y *int) int {
	var temp int
	temp = *x
	*x = *y
	*y = temp

	return temp
}

func main() {

	var a int = 10
	var b int = 20

	fmt.Printf("before a = %d and b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after a = %d and b = %d", a, b)
}
