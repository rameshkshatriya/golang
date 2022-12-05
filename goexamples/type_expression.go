package main

import "fmt"

func main() {
	var value interface{} = "ramesh"
	switch q := value.(type) {
	case bool:
		fmt.Println("value of type is bool")
	case int:
		fmt.Println("value of type is int")
	case string:
		fmt.Println("value of type is string")
	default:
		fmt.Println("value of type %T", q)
	}
}
