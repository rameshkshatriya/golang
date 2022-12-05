package main

import "fmt"

func main() {

	var arr = []string{"hello", "ramesh", "good", "morning", "have", "a", "nice", "day"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}

	for index, element := range arr {
		fmt.Printf("\nindex number: %d,element: %s", index+1, element)
	}

	for _, element := range arr {
		fmt.Printf("\nelement: %s", element)
	}

	var myarr []string
	fmt.Printf("\nlength of string: %d", len(myarr))
	fmt.Printf("\ncapacity of string: %d", cap(myarr))

}
