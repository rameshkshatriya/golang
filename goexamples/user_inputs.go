package main

import "fmt"

func main() {

	var age int
	var name string
	var decimal float32
	var condition bool

	fmt.Println("enter a name, age, temperature, bool: ")
	fmt.Scanln(&name, &age, &decimal, &condition)
	fmt.Printf("name: %s \n age: %d \n temperature: %g \n condition: %t", name, age, decimal, condition)

	// fmt.Println("enter your name: ")
	// fmt.Scanf("%s", &name)

	// fmt.Println("enter a decimal: ")
	// fmt.Scanf("%g", &decimal)

	// fmt.Println("enter true or false: ")
	// fmt.Scanf("%t", &condition)

}
