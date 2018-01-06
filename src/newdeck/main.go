package main

func main() {
	list := newDeck()
	//list.print()
	//hand, remainCards := deal(list, 5)
	//hand.print()
	//fmt.Println()
	//remainCards.print()
	//list.toString()
	list.savetoFile("my_cards")
}
