package main

import "fmt"

func main() {
	var a = make(map[string]string)
	var b map[string]string

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a == nil)
	fmt.Println(b == nil)

}
