package main

import "fmt"

func GFG(i func(p, q string) string) {
	fmt.Println(i("greeks ", "for "))
}

func gfg() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "greeksforgeeks"
	}
	return myf
}

func main() {
	value := func(p, q string) string {
		return p + q + "greeks"
	}
	GFG(value)

	values := gfg()
	fmt.Println(values("welcome ", "to "))

}
