package main

import "fmt"

const value = "ramesh"
const untyped = 124

func main() {

	const num = 55
	value := "suresh"
	const typed int = 125

	fmt.Printf("I got second class in my sslc with :%d\n", num)

	fmt.Printf("my name is %s"+" i got %d in sslc\n", value, num)
	fmt.Println("my register number", untyped)
	fmt.Println("suresh register number is", typed)

}
