package main

import "fmt"

func myfun(a [6]int, size int) int {
	var val, r int
	for i := 0; i < size; i++ {
		val += a[i]
	}
	r = val / size
	return r

}

func main() {
	arr := [6]int{67, 59, 29, 35, 4, 34}

	res := myfun(arr, 6)
	fmt.Printf("final result: %d", res)
}
