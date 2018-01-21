package main

import (
	"fmt"
)

type shape interface {
	getArea() (string, float64)
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {

	sq := square{}
	tri := triangle{}

	sq.sideLength = 10.0
	tri.base = 5.0
	tri.height = 6.0

	printArea(sq)
	printArea(tri)
}

func printArea(ob shape) {
	entity, area := ob.getArea()
	fmt.Println("The area of ", entity, "is", area)
}

func (s square) getArea() (string, float64) {
	return "Square", s.sideLength * s.sideLength
}

func (t triangle) getArea() (string, float64) {
	return "triangle", t.base * t.height
}