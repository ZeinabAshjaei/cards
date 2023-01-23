package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Diamonds", "Hearts", "Spades", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toJsonString() string {
	e, _ := json.Marshal(d)
	return string(e)
}

func (d deck) toJoinString() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(fileName string) error {
	//return os.WriteFile(fileName, []byte(d.toJoinString()), 0666)
	return os.WriteFile(fileName, []byte(d.toJsonString()), 0666)
}

func readFromFile(fileName string) (deck, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(bytes))
	//a := strings.Split(string(bytes), ",")

	name := deck{}
	err2 := json.Unmarshal(bytes, &name)

	return name, err2
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func stringSlicesEqual(a, b deck) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
