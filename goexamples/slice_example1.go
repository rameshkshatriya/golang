package main

import "fmt"

func main() {

	var arr = [8]string{"hello", "ramesh", "good", "morning", "have", "a", "nice", "day"}

	var slice_1 = arr[1:3]
	slice_2 := arr[0:]
	slice_3 := arr[:4]
	slice_4 := arr[:]

	fmt.Println("result slice 1:", slice_1)
	fmt.Println("result slice 2:", slice_2)
	fmt.Println("result slice 3:", slice_3)
	fmt.Println("result slice 4:", slice_4)

	arr1 := []int{10, 20, 30, 40, 50, 60, 70}

	var myslice_1 = arr1[1:3]
	myslice_2 := arr1[0:]
	myslice_3 := arr1[:4]
	myslice_4 := arr1[:]
	myslice_5 := arr1[4:5]

	fmt.Println("result slice 1:", myslice_1)
	fmt.Println("result slice 2:", myslice_2)
	fmt.Println("result slice 3:", myslice_3)
	fmt.Println("result slice 4:", myslice_4)
	fmt.Println("result slice 5:", myslice_5)

	var myarr = make([]int, 3, 5)
	fmt.Printf("slice: %v,\nlength: %d,\ncapacity: %d", myarr, len(myarr), cap(myarr))

	myarr1 := make([]int, 7)
	fmt.Printf("\nslice: %v,\nlength: %d,\ncapacity: %d", myarr1, len(myarr1), cap(myarr1))

}
