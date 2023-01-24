package main

import (
	"fmt"
	"github.com/ZeinabAshjaei/cards/deckstruct"
	"os"
)

func main() {
	cards := deckstruct.NewDeck()
	dealCards, remainingCards := deckstruct.Deal(cards, 5)
	dealCards.Print()
	fmt.Println("###################")
	remainingCards.Print()

	//a := cards.toJoinString()
	//fmt.Println(a)

	var err error

	err = cards.SaveToFile("my_cards_struct")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	a, err := deckstruct.ReadFromFile("my_cards_struct")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	a.Print()

	cards.Shuffle()
	fmt.Println("###################")
	cards.Print()
}
