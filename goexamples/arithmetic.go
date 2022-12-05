package main

import "fmt"

func main() {

	a := 30
	b := 18

	result1 := a + b
	fmt.Printf("adding a + b = %d\n", result1)
	fmt.Printf("type a + b = %T", result1)

	result2 := a - b
	fmt.Println("\nsubtract a - b = ", result2)

	result3 := a * b
	fmt.Printf("multiply a * b = %d", result3)

	result4 := a / b
	fmt.Printf("\ndivide a / b = %d", result4)

	result5 := a % b
	fmt.Printf("\nmodulus a %% b = %d", result5)
}
