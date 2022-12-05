// Go program to illustrate how to
// access the bytes of the string
package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a string
	str := "Hello Ramesh"

	// Accessing the bytes of the given string
	for c := 0; c < len(str); c++ {

		fmt.Printf("\nCharacter = %c Bytes = %v", str[c], str[c])
	}

	str1 := []byte{0x58, 0x41, 0x53, 0x45, 0x59, 0x48}
	st := string(str1)
	fmt.Println("str1: ", st)
}
