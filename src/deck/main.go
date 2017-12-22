package main

type deck []string

func main() {
	cards := deck{"King of spades", "Heart of Ace", newCard(), "Joker"}
	cards.print()
}

func newCard() string {
	return "Queen of Club"
}
