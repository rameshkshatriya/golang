package main

import "fmt"

func main() {

	var arr = [8]string{"hello", "ramesh", "good", "morning", "have", "a", "nice", "day"}

	fmt.Println("arr:", arr)

	var slice = arr[1:4]
	fmt.Println("slice:", slice)

	fmt.Printf("length of slice: %d", len(slice))

	fmt.Printf("\ncapacity of slice: %d\n", cap(slice))

	var arr1 = []string{"moin", "megu", "madhu", "prabhu"}
	fmt.Println("print arr1:", arr1)

	arr2 := []string{"abcd", "efgh", "ijkl", "mnop"}
	fmt.Println("print arr2:", arr2)
}
