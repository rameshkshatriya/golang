package main

import "fmt"

func fun(n int) {
	if n > 0 {
		k := n * *2
		fmt.Println(k)
		fun(n - 1)
	}

}

func main() {
	fun(4)
}
