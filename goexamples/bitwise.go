package main

import "fmt"

func main() {

	a := 30
	b := 14

	result1 := a & b
	fmt.Printf("result1 of a & b = %d", result1)

	result2 := a | b
	fmt.Printf("\nresult2 of a | b = %d", result2)

	result3 := a ^ b
	fmt.Printf("\nresult3 of a ^ b = %d", result3)

	result4 := a << 1
	fmt.Printf("\nresult4 of a << 1 = %d", result4)

	result5 := a >> 1
	fmt.Printf("\nresult5 of a >> 1 = %d", result5)

	result6 := a &^ b
	fmt.Printf("\nresult6 of a &^ b = %d", result6)

}
