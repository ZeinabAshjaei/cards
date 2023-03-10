package decktypestring

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := NewDeck()

	if len(cards) != 16 {
		t.Errorf("The length of deck is wrong, it is %v", len(cards))
	}

	if cards[0] != "Ace of Diamonds" {
		t.Errorf("The first card is wrong, it is %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Club" {
		t.Errorf("The last card is wrong, it is %v", cards[len(cards)-1])
	}
}

func TestToJsonString(t *testing.T) {
	cards := deck{"a", "b", "c"}
	jsonCards := cards.ToJsonString()

	expected := fmt.Sprintf("[%q,%q,%q]", "a", "b", "c")
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

	var fromFile deck
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
