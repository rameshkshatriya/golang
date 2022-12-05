package main

import "fmt"

type class struct {
	name      string
	language  string
	gopercent float64
}

func main() {

	result := class{name: "moin", language: "golang", gopercent: 98.99}

	fmt.Println("class name:", result.name)
	fmt.Println("class language:", result.language)
	fmt.Println("class gopercent:", result.gopercent)

}
