package main

import (
	"fmt"
)

type Animal interface {
	predator()
}

type lion struct{
	color string
	height float32
	length	float32
}

type tiger struct {
	color string
	height float32
	length float32
}

func main() {

	l := lion{"Yellow", 4.7, 7.2}
	t := tiger{"White", 4.67, 7.2}

	attribute(l)
	attribute(t)
}

func attribute(a Animal) {
	a.predator()
}

func (l lion) predator() {
	fmt.Println(l.color, l.height, l.length)
}

func (t tiger) predator() {
	fmt.Println(t.color, t.height, t.length)
}