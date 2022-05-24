package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	card_suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	card_values := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range card_suits {
		for _, value := range card_values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, hadn_size int) (deck, deck) {
	return d[:hadn_size], d[hadn_size:]
}

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
