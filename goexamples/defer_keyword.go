package main

import "fmt"

func mul(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("result:", res)
	return 0

}
func show() {
	fmt.Println("hello ramesh")
}

func main() {

	fmt.Println("start")
	defer fmt.Println("end")
	mul(10, 20)
	defer mul(20, 30)
	show()

}
