package main

import (
	"fmt"
	"strings"
	"os"
	"math/rand"
	"io/ioutil"
)


type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards 
}

func deal(cards deck, number int) (deck, deck)  {

	return cards[:number], cards[number:]
	
}

func (d deck) print() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(bs)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle(){
	for i := range d{
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}