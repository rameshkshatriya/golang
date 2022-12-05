package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := []byte{'!', '!', 'g', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'g', 'e', 'e', 'k', 's', '@', '@'}

	res1 := bytes.Split(s1, []byte("eek"))

	fmt.Printf("%s\n", res1)

	s2 := bytes.Split([]byte("!!geeksforgeeks@@"), []byte("eek"))

	fmt.Printf("%s", s2)

}
