package main

import "fmt"

func main() {
	s1 := "Hello"
	s2 := "hello"
	s3 := "Hello"
	result1 := s1 == s2
	result2 := s1 == s3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("the type od result1 is %T and "+"the type of result2 is %T", result1, result2)
}
