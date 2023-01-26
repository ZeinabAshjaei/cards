package deckstruct

import (
	"encoding/json"
	"fmt"
	"github.com/ZeinabAshjaei/cards/filehandling"
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Cards       []string
	fileService filehandling.FileManager
}

func NewDeck(FileService filehandling.FileManager) *Deck {
	deckCards := make([]string, 0)

	cardSuites := []string{"Diamonds", "Hearts", "Spades", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			deckCards = append(deckCards, value+" of "+suite)
		}
	}

	d := Deck{
		Cards:       deckCards,
		fileService: FileService,
	}
	return &d
}

func (d *Deck) Print() {
	for i, card := range d.Cards {
		fmt.Println(i, card)
	}
}

func Deal(d *Deck, handSize int) (*Deck, *Deck) {
	dealCards := Deck{Cards: d.Cards[:handSize]}
	remainingCards := Deck{Cards: d.Cards[handSize:]}

	return &dealCards, &remainingCards
}

func (d *Deck) ToJsonString() string {
	e, _ := json.Marshal(d)
	return string(e)
}

func (d *Deck) ToJoinString() string {
	return strings.Join(d.Cards, ", ")
}

func (d *Deck) SaveToFile(fileName string) error {
	//return os.WriteFile(fileName, []byte(d.ToJoinString()), 0666)
	return d.fileService.WriteFile(fileName, []byte(d.ToJsonString()), 0666)
}

func (d *Deck) ReadFromFile(fileName string) error {
	fileBytes, err := d.fileService.ReadFile(fileName)

	if err != nil {
		return err
	}

	err = json.Unmarshal(fileBytes, d)
	if err != nil {
		fmt.Println("Cannot convert string representation to deck")
	}

	return nil
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func StringSlicesEqual(a, b *Deck) bool {
	if len(a.Cards) != len(b.Cards) {
		return false
	}
	for i, v := range a.Cards {
		if v != b.Cards[i] {
			return false
		}
	}
	return true
}
