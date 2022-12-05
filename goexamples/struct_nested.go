package main

import "fmt"

type student struct {
	name    string
	class   string
	subject int
}

type teacher struct {
	name    string
	subject string
	exp     int
	details student
}

func main() {

	result := teacher{
		name:    "moin",
		subject: "programming",
		exp:     6,
		details: student{name: "ramesh", class: "python", subject: 3},
	}

	fmt.Println("Details of teacher")
	fmt.Println("teacher name:", result.name)
	fmt.Println("teacher subject:", result.subject)
	fmt.Println("teacher exp:", result.exp)

	fmt.Println("\nDetails of students")
	fmt.Println("student name:", result.details.name)
	fmt.Println("student class:", result.details.class)
	fmt.Println("student subject:", result.details.subject)

}
