package main

import "fmt"

func main() {
	var val, r int
	a := [6]int{67, 59, 29, 35, 4, 34}
	for i := 0; i < 6; i++ {
		val += a[i]
		fmt.Println(val)

	}
	r = val / 6
	fmt.Printf("final ans:%d", r)
}
