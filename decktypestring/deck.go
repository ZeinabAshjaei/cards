package decktypestring

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func NewDeck() deck {
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

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToJsonString() string {
	e, _ := json.Marshal(d)
	return string(e)
}

func (d deck) ToJoinString() string {
	return strings.Join(d, ", ")
}

func (d deck) SaveToFile(fileName string) error {
	//return os.WriteFile(fileName, []byte(d.ToJoinString()), 0666)
	return os.WriteFile(fileName, []byte(d.ToJsonString()), 0666)
}

func ReadFromFile(fileName string) (deck, error) {
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

func (d deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func StringSlicesEqual(a, b deck) bool {
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
