package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Make a new type of deck
// its a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, suit+" of "+number)
		}
	}
	return cards
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromDisk(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	d := strings.Split(string(bs), ",")
	return deck(d)
}

func (d deck) shuffel() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d {
		random := r.Intn(len(d) - 1)
		d[i], d[random] = d[random], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
