package main

import "fmt"

func main() {
	var car = map[string]string{"brand": "Rolls-royce", "model": "goast", "year": "2020", "day": "", "pric1": "xxxxxxxxxx"}

	val1, ok1 := car["brand"]
	val2, ok2 := car["model"]
	val3, ok3 := car["day"]
	val4, ok4 := car["prices"]
	_, ok5 := car["yeas"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(val4, ok4)
	fmt.Println(ok5)
}
