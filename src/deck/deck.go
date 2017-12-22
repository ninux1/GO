package main

import "fmt"

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
