package main

import "fmt"

type author struct {
	name   string
	branch string
}

func (a *author) show_1(abranch string, aname string) {
	(*a).branch = abranch
	(*a).name = aname

}

func (a author) show_2() {
	a.name = "suresh"
	fmt.Println("before author name: ", a.name)
}

func main() {

	res := author{
		name:   "ramesh",
		branch: "mech",
	}

	fmt.Println("before author branch: ", res.branch)
	fmt.Println("before author name: ", res.name)

	res.show_1("computer", "karthik")

	fmt.Println("after author branch: ", res.branch)
	fmt.Println("after author name: ", res.name)

	(&res).show_2()
	fmt.Println("after author name: ", res.name)

}
