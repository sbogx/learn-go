package main

import ()

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
