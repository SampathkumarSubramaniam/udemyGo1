package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for _, value := range d {
		fmt.Println(value)
	}
}

func (d deck) hello() {
	fmt.Println("In am available only for deck type.")
}

func  newDeck() deck {
	card := deck{}
	cardPrefix := []string{"Ace", "One", "Two", "Three"}
	cardSuffix := []string{"Diamond", "Spades", "someThingElse"}
	for _, preValue := range cardPrefix {
		for _, sufValue := range cardSuffix {
			card = append(card, preValue+"-"+sufValue)
		}
	}
	return card
}

func dealDeck(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error occured while reading file:", filename, "Error:", err)
		os.Exit(1)
	}

	abc := string(bs)
	abc1 := strings.Split(abc, ",")
	return deck(abc1)

}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
