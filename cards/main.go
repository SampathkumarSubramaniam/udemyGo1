package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffleDeck()
	fmt.Println("---------------")
	cards.print()
	fmt.Println("---------------")
	/*hand, remainingDeck := dealDeck(cards, 8)
	hand.print()
	remainingDeck.print()
	fmt.Println("--------------------")
	fmt.Println(cards.toString())
	fmt.Println("--------------------")*/
	fmt.Println(cards.saveToFile("temp.txt"))
}
