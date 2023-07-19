package main

import (
	"fmt"

	card "github.com/AishwaryaKhonde/cards/card"
)

func main() {
	fmt.Println("hello")
	//var card string = "Cards"
	//card := "Cards"
	//card := []string{"one of spades", "one of hearts", NewDeck()}
	//card := deck{"one of spades", "one of hearts"}

	cards := card.NewDeck()
	cards.Print()
	cards.SaveToFile("newCrad")
	selectedCard := cards.PickOne()
	fmt.Println("selectedCard : ", selectedCard)
	//fmt.Println(card)
	// for _, values := range card {
	// 	fmt.Println(values)

	// }
}

// func NewDeck() string {
// 	return "one of diamonds"
// }
