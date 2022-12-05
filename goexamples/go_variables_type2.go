package main

import "fmt"

func main() {
	myvar1 := 24
	myvar2 := "helloramesh"
	myvar3 := 24.25
	mynum1, mynum2, mynum3 := 10, 20, 30
	my1, my2 := 39, 200
	my3, my2 := 45, 100
	my_1, my_2, my_3 := 124, "hello", 124.23

	fmt.Printf("the value of myvar1 is : %d\n", myvar1)
	fmt.Printf("the type of myvar1 is : %T\n", myvar1)
	fmt.Printf("the value of myvar2 is : %s\n", myvar2)
	fmt.Printf("the type of myvar2 is : %T\n", myvar2)
	fmt.Printf("the value of myvar3 is : %f\n", myvar3)
	fmt.Printf("the type of myvar3 is : %T\n", myvar3)
	fmt.Printf("the value of mynum1 is : %d\n", mynum1)
	fmt.Printf("the value of mynum2 is : %d\n", mynum2)
	fmt.Printf("the value of mynum3 is : %d\n", mynum3)
	fmt.Printf("the value of my1_my2 is : %d %d\n", my1, my2)
	fmt.Printf("the value of my3_my2 is : %d %d\n", my3, my2)
	fmt.Printf("the value of my_1 is : %d\n", my_1)
	fmt.Printf("the value of my_2 is : %s\n", my_2)
	fmt.Printf("the value of my_3 is : %f\n", my_3)
}
