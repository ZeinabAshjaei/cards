package main

func main() {
	cards := newDeck()
	// dealCards, remainingCards := deal(cards, 5)
	// dealCards.print()
	// remainingCards.print()

	//a := cards.toJoinString()
	//fmt.Println(a)
	//var err error
	//var a deck

	/*err = cards.saveToFile("my_cards")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	a, err = readFromFile("my_cards")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	a.print()*/

	cards.shuffle()
	cards.print()
}
