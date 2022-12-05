package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		fmt.Println("hello ramesh")
	}
	j := 0
	for j < 3 {
		j += 2
	}
	fmt.Println(j)

	variable := []string{"hello", "ramesh", "good_morning"}

	for a, b := range variable {
		fmt.Println(a, b)
	}

	integer := []int{1001, 1002, 1003, 1004}
	for c, d := range integer {
		fmt.Println(c, d)
	}

	decimal := []float64{10.1, 10.2, 10.3}
	for _, f := range decimal {
		fmt.Println(f)
	}

	for x, y := range "Xabcd" {
		fmt.Printf("the index number of %U is %d\n", y, x)
	}

	for p, q := range "ramesh" {
		fmt.Printf("the index of %d is %c\n", p, q)
	}

	mmap := map[int]string{
		2984: "moin",
		2987: "meghanath",
		3076: "madhusudhan",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	for key, value := range mmap {
		fmt.Printf("%s working in msys technologies, emp no:%d\n", value, key)
	}

	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}

	queue := make(chan string, 3)
	queue <- "moin"
	queue <- "meghanath"
	queue <- "madhusudhan"
	close(queue)
	for v := range queue {
		fmt.Println(v)
	}

}
