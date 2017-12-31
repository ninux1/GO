package main

import "fmt"

func main() {
	list := newDeck()
	//list.print()
	hand, remainCards := deal(list, 5)
	hand.print()
	fmt.Println()
	remainCards.print()
}
