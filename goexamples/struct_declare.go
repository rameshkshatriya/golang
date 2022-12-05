package main

import "fmt"

type address struct {
	slno     int
	name     string
	city     string
	landmark string
	pincode  int
}

func main() {

	var a address
	fmt.Println(a)

	a1 := address{01, "ramesh", "hebbal", "near gas egency", 560024}
	fmt.Println("address1: ", a1)

	a2 := address{slno: 02, name: "suresh", pincode: 560009, landmark: "gas pipe", city: "tabala"}
	fmt.Println("address2 :", a2)

	a3 := address{pincode: 500342}
	fmt.Println("address:", a3)

}
