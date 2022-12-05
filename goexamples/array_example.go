package main

import "fmt"

func main() {

	myarr := [3][4]string{{"moin", "megu", "ramesh"}, {"vishal", "vinayak", "suhas", "nandu"}, {"teacher", "karthik", "chandh", "manu"}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(myarr[i][j])
		}
	}

	var array [2][2]int
	array[0][0] = 100
	array[0][1] = 200
	array[1][0] = 300
	array[1][1] = 400

	for k := 0; k < 2; k++ {
		for l := 0; l < 2; l++ {
			fmt.Println(array[k][l])
		}

	}

	var myarr1 [3][4]string
	myarr1[0][0] = "moin"
	myarr1[0][1] = "megu"
	myarr1[0][2] = "ramesh"
	myarr1[1][0] = "vishal"
	myarr1[1][1] = "vinayak"
	myarr1[1][2] = "suhas"
	myarr1[1][3] = "nandhu"
	myarr1[2][0] = "teacher"
	myarr1[2][1] = "karthik"
	myarr1[2][2] = "chandh"
	myarr1[2][3] = "manu"

	for m := 0; m < 3; m++ {
		for n := 0; n < 4; n++ {
			fmt.Println(myarr1[m][n])
		}

	}

	var myint [2]int
	fmt.Println("elements of the array:", myint)

	var mystr [2]string
	fmt.Println("elements of the array:", mystr)

}
