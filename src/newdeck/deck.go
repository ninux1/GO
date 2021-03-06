package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) savetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
