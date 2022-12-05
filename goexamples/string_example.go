package main

import "fmt"

func main() {

	s1 := "hello good morning"
	fmt.Println(s1)

	var s2 string
	s2 = "have a nice day"
	fmt.Println(s2)

	s3 := `Hello!GeeksforGeeks`
	fmt.Println(s3)

	str1 := "hellogoodmorning"
	fmt.Println(str1)

	for index, chr := range "helloramesh" {
		fmt.Printf("the index of %c is %d\n", chr, index)

	}

	str := "Welcome to GeeksforGeeks"

	for c := 0; c < len(str); c++ {

		fmt.Printf("\ncharacter = %c bytes = %v", str[c], str[c])

	}

}
