package main

import "fmt"

func main() {

	arr1 := []int{10, 20, 30}
	arr2 := [...]int{10, 20, 30, 40}
	arr3 := [3]int{10, 20, 30}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	fmt.Println(len(arr3))

	myvar := [...]int{20, 30, 40, 50, 60, 70, 80, 90}

	for i := 0; i < len(myvar); i++ {

		fmt.Printf("%d\n", myvar[i])
	}

	o_array := [...]int{100, 200, 300, 400, 500}
	fmt.Println("original array:", o_array)

	new_array := o_array
	fmt.Println("new array:", new_array)

	new_array[0] = 500
	fmt.Println("new array(after changes):", new_array)

	fmt.Println("original array{any changes:", o_array)

	myarr1 := [3]int{10, 20, 30}
	myarr2 := [...]int{10, 20, 30}
	myarr3 := [3]int{10, 30, 40}

	fmt.Println(myarr1 == myarr2)
	fmt.Println(myarr1 == myarr3)
	fmt.Println(myarr2 == myarr3)

}
