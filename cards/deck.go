package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create of a new type of type Deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsNumbers := []string{"Ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	for _, Suite := range cardsSuits {
		for _, number := range cardsNumbers {
			cards = append(cards, number+" of "+Suite)
		}
	}
	return cards
}

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) saveTofile(filename string) error {
	return os.WriteFile(filename, []byte(cards.toString()), 0666)
}

func newdeckfromfile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error occured ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rand_ := rand.New(source)
	for i := range d {
		random_number := rand_.Intn(len(d) - 1)
		d[i], d[random_number] = d[random_number], d[i]
	}
}
