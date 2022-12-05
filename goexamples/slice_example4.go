package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := [][]int{{10, 20},
		{30, 40},
		{40, 50},
		{60, 70},
	}
	fmt.Println(arr)

	arr1 := [][]string{
		[]string{"hello", "ramesh"},
		[]string{"good", "morning"},
		[]string{"nice", "day"},
	}

	fmt.Println(arr1)

	fmt.Println("sorting")

	myarr1 := []int{20, 10, 9, 15, 25, 4, 8, 30}
	myarr2 := []string{"python", "golang", "scala", "c#", "java", "html"}

	fmt.Println("original myarr1:", myarr1)
	fmt.Println("original myarr2:", myarr2)

	sort.Ints(myarr1)
	sort.Strings(myarr2)

	fmt.Println("original myarr1:", myarr1)
	fmt.Println("original myarr2:", myarr2)

}
