package main

import (
	"fmt"
	"strings"
)

func strjoin(elements ...string) string {
	return strings.Join(elements, "-")
}

func main() {
	fmt.Printf(strjoin())

	fmt.Printf(strjoin("greek", "for", "greks\n"))

	element := []string{"hello", "ramesh", "good morning"}
	fmt.Printf(strjoin(element...))
}
