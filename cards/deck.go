package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create new type of deck, simillar to a slice of string. It is same as creating dynamic array with user defined functionality

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Ten", "Nine", "Eight", "Seven", "Six", "Five", "Four", "Three", "Two"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

// example of how to return multiple values from a function in GO
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// value := []string(d)
	return strings.Join([]string(d), ",\n")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		//generate a random number within length
		newPosition := r.Intn(len(d) - 1)

		// swap the position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
