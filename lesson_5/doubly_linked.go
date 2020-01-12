package main

import "fmt"

type element struct {
	name string
	age  int
}

type list struct {
	elem []element
	len  int
}

func main() {

	a := element{"SM", 13}
	g := element{"AZ", 20}
	b := list{}

	fmt.Println(b.Add(&a))
	fmt.Println(b.Add(&g))
	fmt.Println(b.elem)
	fmt.Println(b.len)
}

func (l *list) Add(c *element) *list {
	l.elem = append(l.elem, *c)
	l.len++
	return l
}
