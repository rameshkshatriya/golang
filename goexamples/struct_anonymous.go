package main

import "fmt"

func main() {

	element := struct {
		name    string
		city    string
		pincode int
	}{
		name:    "moin",
		city:    "yeshwanthpur",
		pincode: 560029,
	}

	fmt.Println(element)

}
