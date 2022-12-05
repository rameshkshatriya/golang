package main

import "fmt"

type car struct {
	name, model, color string
	make               int
}

func main() {

	c := car{name: "lamborgini", model: "urus", color: "yellow", make: 2021}

	fmt.Println("car name: ", c.name)

	fmt.Println("car color: ", c.color)

	c.color = "black"

	fmt.Println("car color: ", c.color)

	fmt.Println("full details of car :", c)

	c.make = 2022

	fmt.Println("full details of car :", c)

	c.name = "BMW"

	fmt.Println("full details of car :", c)

	c.model = "720d"

	fmt.Println("full details of car :", c)

}
