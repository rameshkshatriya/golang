package main

import (
	"fmt"
)

const pi = 3.14

func main() {

	const trueconst = true

	type mybool bool

	var defaultbool = trueconst
	var custombool mybool = trueconst

	fmt.Println(defaultbool)
	fmt.Println(custombool)
}
