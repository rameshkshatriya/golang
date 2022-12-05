package main

import "fmt"

func main() {

	arr := []int{22, 33, 44, 55, 66, 77}
	arr1 := arr[0:4]
	fmt.Println("original arr:", arr)
	fmt.Println("slice result:", arr1)

	fmt.Println("modified")

	arr1[0] = 222
	arr1[1] = 333
	arr1[2] = 444

	fmt.Println("original arr:", arr)
	fmt.Println("slice result:", arr1)

	fmt.Println("comparision of slice")

	s1 := []int{10, 20, 30}
	var s2 []int

	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

}
