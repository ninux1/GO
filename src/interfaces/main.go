package main

import "fmt"

type bot interface {
	Greet() string
}

type ebot struct{} 
type sbot struct{}

func main() {
	eb := ebot{}
	sb := sbot{}

	printGreet(eb)
	printGreet(sb)
}

func printGreet(b bot) {
	fmt.Println(b.Greet())
}

func (ebot) Greet() string {
	return "Hi there"
}

func (sbot) Greet() string {
	return "Hola"
}
