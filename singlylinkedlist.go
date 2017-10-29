package main

import "fmt"

type list struct {
	front *node
}

type node struct{
	next *node
	data interface{}
}

type List interface {

	insert(data interface{})
	display()
}


func (d * node) init () *node {
	d.next = nil
	d.data = nil
	return d
}

func Newnode() *node {
	 return new(node).init()
}
func (d* list) init() *list {
	d.front = nil
	return d
}

func NewList() *list {
	return new(list).init()
}

func (s *list) insert(data *node) {

	Newnode()
	if s.front == nil {
		s.front = data
	} else {

		var temp *node
		temp = s.front
		for temp.next != nil {
		temp = temp.next
		}

		temp.next = data
	}

}

func (s *list)display() {
temp := s.front
	for temp != nil {
		fmt.Println("Data %v",temp.data)
		temp = temp.next
	}

}

func main() {
	s := NewList()
	n1 := &node{data: "test1"}
	s.insert(n1)
	n2 := &node{data: "test2"}
	s.insert(n2)

	s.display()

}

