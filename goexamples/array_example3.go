package main

import "fmt"

func main() {

	arr1 := [...]string{"apple", "ball", "cat", "dog"}
	arr2 := arr1

	fmt.Println("elements of arr1:", arr1)
	fmt.Println("elements of arr2:", arr2)

	arr1[0] = "amma"

	fmt.Println("elements of arr1:", arr1)
	fmt.Println("elements of arr2:", arr2)

	myarr1 := [...]int{10, 20, 30, 40, 50}
	myarr2 := &myarr1

	fmt.Println("elements of myarr1:", myarr1)
	fmt.Println("elements of myarr2:", *myarr2)

	myarr1[4] = 500

	fmt.Println("elements of myarr1:", myarr1)
	fmt.Println("elements of myarr2:", *myarr2)

}
