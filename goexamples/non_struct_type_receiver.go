package main

import "fmt"

type data int

func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

func main() {

	value1 := data(3)
	value2 := data(2)
	res := value1.multiply(value2)
	fmt.Println("final result: ", res)
}
