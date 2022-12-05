package main

import "fmt"

var num int

func prime_or_not(value int) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println("not prime")
			break
		} else {
			fmt.Println("prime")
		}

	}
}

// func main() {
// 	fmt.Println("prime numbers")
// 	var result string = prime_or_not(3)
// 	fmt.Println(result)
// }
