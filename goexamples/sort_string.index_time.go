package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {

	s := []int{10, 8, 2, 4, 16, 7, 29, 12}
	fmt.Println("unsorted array: ", s)
	sort.Ints(s)
	fmt.Println("sorted array: ", s)

	// str := []string{"greeksforgeek"}
	fmt.Println("index value: ", strings.Index("greeksforfeek", "fo"))
	// fmt.Println("index vaue: ", str)

	fmt.Println("time: ", time.Now().Unix())
}
