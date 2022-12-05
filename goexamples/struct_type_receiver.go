package main

import "fmt"

type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

func (a author) show() {
	fmt.Println("author name: ", a.name)
	fmt.Println("author branch: ", a.branch)
	fmt.Println("author particles: ", a.particles)
	fmt.Println("author salary: ", a.salary)
}

func main() {

	res := author{
		name:      "ramesh",
		branch:    "mech engg",
		particles: 200,
		salary:    400000,
	}
	res.show()
}
