package main

import "fmt"

func main()  {
	cards := newDeck()
	fmt.Println(cards.toString())
}

func newCard() string{
	return "Five of Spade"
}