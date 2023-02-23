package main

func main() {
	//cards := newDeck()

	//fmt.Println(cards.toString())

	//cards.saveToFile("my_deck")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
