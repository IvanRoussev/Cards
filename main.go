package main

func main() {
	cards := newDeck()
	cards.newDeckFromFile("my_dk")
	cards.print()
}
