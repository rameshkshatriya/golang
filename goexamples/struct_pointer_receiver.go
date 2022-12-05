package main

import "fmt"

type author struct {
	name    string
	branch  string
	subject string
}

func (a *author) show(aname string, abranch string, asubject string) {
	(*a).branch = abranch
	(*a).name = aname
	(*a).subject = asubject

}

func main() {

	res := author{
		name:    "RAMESH",
		branch:  "ME",
		subject: "AUTOMOBILE",
	}

	fmt.Println("author name: ", res.name)
	fmt.Println("author before branch: ", res.branch)
	fmt.Println("author subject: ", res.subject)

	p := (&res)
	p.show("SURESH", "CSE", "PROGRAMMING")

	fmt.Println("author before name: ", res.name)
	fmt.Println("author after branch: ", res.branch)
	fmt.Println("author subject: ", res.subject)

}
