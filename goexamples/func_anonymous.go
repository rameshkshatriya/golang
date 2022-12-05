package main

import "fmt"

func main() {

	func() {
		fmt.Println("hello ramesh good morning")
	}()

	value := func() {
		fmt.Println("good morning")
	}
	value()

	func(ele string) {
		fmt.Println(ele)
	}("helo moin")

}
