package deckstruct

import (
	"fmt"
	"github.com/ZeinabAshjaei/cards/filehandling"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	//fs := filehandling.MockFileService{}
	cards := NewDeck(nil)

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

	expectedDeck := Deck{
		Cards:       []string{"a", "b", "c"},
		fileService: filehandling.NewMockFileService(),
	}

	err = expectedDeck.SaveToFile(testFileName)
	if err != nil {
		t.Errorf("save to file failed for error %v", err)
	}

	actualDeck := &Deck{
		fileService: filehandling.NewMockFileService(),
	}

	err = actualDeck.ReadFromFile(testFileName)
	if err != nil {
		t.Errorf("Cannot read Deck from file")
	}

	if len(actualDeck.Cards) != 3 ||
		actualDeck.Cards[0] != expectedDeck.Cards[0] ||
		actualDeck.Cards[2] != expectedDeck.Cards[2] {
		t.Errorf("Cannot read Deck from file")
	}

	err = os.Remove(testFileName)
	if err != nil && !os.IsNotExist(err) {
		t.Errorf("Error while deleting test file %v", err)
	}
}
