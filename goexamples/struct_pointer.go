package main

import "fmt"

type employee struct {
	name, city, salary string
	age                int
}

func main() {

	emp1 := &employee{name: "moin", city: "yeshwanthpur", age: 23, salary: "4lpa"}

	fmt.Println("emp name:", (*emp1).name)

	fmt.Println("emp city:", emp1.city)
}
