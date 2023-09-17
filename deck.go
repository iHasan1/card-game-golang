package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}


	//  we add underscore in place of variables we dont plan to use 
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
} 

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string, ) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - LOG THE ERROR AND MAKE A CALL TO NEWDECK()
		// Option #2 - Log the error and entire quit the program
		fmt.Println("Error: ", err) 
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
