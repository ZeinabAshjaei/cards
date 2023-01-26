package main

import (
	"fmt"
	"github.com/ZeinabAshjaei/cards/deckstruct"
	"github.com/ZeinabAshjaei/cards/filehandling"
	"os"
)

func main() {
	cards := deckstruct.NewDeck(filehandling.NewFileService())

	/*dealCards, remainingCards := deckstruct.Deal(cards, 5)
	dealCards.Print()
	fmt.Println("###################")
	remainingCards.Print() */

	//a := cards.toJoinString()
	//fmt.Println(a)

	var err error

	err = cards.SaveToFile("my_cards_struct")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cards.ReadFromFile("my_cards_struct")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cards.Print()

	cards.Shuffle()
	fmt.Println("###################")
	cards.Print()
}
