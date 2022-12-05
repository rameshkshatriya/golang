package main

import "fmt"

type node struct {
	next *node
	data int
}

type link struct {
	head   *node
	length int
}

func (l *link) add_beginning(value int) {
	newNode := node{data: value}
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
	return
}

func (l *link) displaylist() {
	if l.head == nil {
		return
	}
	currentnode := l.head
	for currentnode != nil {
		fmt.Println(currentnode.data)
		currentnode = currentnode.next
	}
}

func main() {
	mylist := link{}
	mylist.add_beginning(10)
	mylist.add_beginning(20)
	mylist.add_beginning(30)
	mylist.add_beginning(40)
	mylist.add_beginning(50)
	mylist.displaylist()

}
