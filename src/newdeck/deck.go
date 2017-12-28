package main

import "fmt"

type deck []string

func newDeck() deck {
	new := deck{}
	category := deck{"Club", "Spade", "Heart", "Diamond"}
	cards := deck{"Ace of ", "Jack of ", "Queen of ", "King of "}

	for _, val1 := range category {
		for _, val2 := range cards {
			new = append(new, val2+val1)
		}
	}
	return new
}

func (d deck) print() {
	for _, elements := range d {
		fmt.Println(elements)
	}
}
