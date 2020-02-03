package main

// import "fmt"

func main()  {
	// cards := newDeckFromFile("my_cards1")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}