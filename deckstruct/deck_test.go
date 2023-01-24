package deckstruct

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := NewDeck()

	if len(cards.Cards) != 16 {
		t.Errorf("The length of Deck is wrong, it is %v", len(cards.Cards))
	}

	if cards.Cards[0] != "Ace of Diamonds" {
		t.Errorf("The first card is wrong, it is %v", cards.Cards[0])
	}

	if cards.Cards[len(cards.Cards)-1] != "Four of Club" {
		t.Errorf("The last card is wrong, it is %v", cards.Cards[len(cards.Cards)-1])
	}
}

func TestToJsonString(t *testing.T) {
	myCards := []string{"a", "b", "c"}
	cards := Deck{Cards: myCards}
	jsonCards := cards.ToJsonString()

	expected := fmt.Sprintf("{\"Cards\":[%q,%q,%q]}", "a", "b", "c")
	if jsonCards != expected {
		t.Errorf("The toJsonString func is not OK")
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	testFileName := "test_file"
	err := os.Remove(testFileName)

	cards := NewDeck()

	err = cards.SaveToFile(testFileName)
	if err != nil {
		t.Errorf("save to file failed for error %v", err)
	}

	var fromFile *Deck
	fromFile, err = ReadFromFile(testFileName)
	if err != nil {
		t.Errorf("read from file failed for error %v", err)
	}

	if !(StringSlicesEqual(fromFile, cards)) {
		t.Errorf("read from file failed for error %v", err)
	}

	err3 := os.Remove(testFileName)
	if err3 != nil {
		t.Errorf("removing file failed for error %v", err)
	}
}
