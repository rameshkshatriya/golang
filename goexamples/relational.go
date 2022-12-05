package main

import "fmt"

func main() {

	a := 34
	b := 40

	result1 := a == b
	fmt.Println("result1 = ", result1)

	result2 := a != b
	fmt.Println("result2 = ", result2)

	result3 := a > b
	fmt.Println("result3 = ", result3)

	result4 := a < b
	fmt.Println("result4 = ", result4)

	result5 := a >= b
	fmt.Println("result5 = ", result5)

	result6 := a <= b
	fmt.Println("result6 = ", result6)
}
