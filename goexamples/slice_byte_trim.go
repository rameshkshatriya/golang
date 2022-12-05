package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := []byte{'!', '!', 'g', 'e', 'e', 'k', 's', '@', '@'}

	res1 := bytes.Trim(s1, "!@")

	fmt.Printf("%s", res1)

	s2 := bytes.Trim([]byte("!!geeks@@for@@geeks##"), "!@@#")

	fmt.Printf("\n%s", s2)

}
