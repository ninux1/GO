package main

import (
	"fmt"
)

type m interface {
	service() int
}

type Truck struct {
	Name string
}

type Car struct {
	Name string
}

func main() {
	T := Truck{"4Runner"}
	C := Car{"Camry"}

	fmt.Println("Charges for Truck =", charges(T))
	fmt.Println("charges for Car = ", charges(C))

}

func charges(i m) int {

	c := i.service()
	return c
}


func (t Truck) service() int {
	if t.Name == "4Runner" {
		fmt.Println("Its a Truck")
	}
	return 50
}

func (c Car) service() int {
	if c.Name == "Camry" {
		fmt.Println("Its a Car")
	}
	return 30
}