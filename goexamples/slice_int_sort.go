package main

import (
	"fmt"
	"sort"
)

func main() {

	s1 := []int{400, 500, 300, 200, 100, 700}
	s2 := []int{-23, -9, 10, -4, 20, -10}

	sort.Ints(s1)
	sort.Ints(s2)

	fmt.Println(s1)
	fmt.Println(s2)

	s3 := []int{100, 200, 300, 400, 500}
	s4 := []int{-23, -9, 10, -4, 20, -10}

	res1 := sort.IntsAreSorted(s3)
	res2 := sort.IntsAreSorted(s4)

	fmt.Println(res1)
	fmt.Println(res2)

}
