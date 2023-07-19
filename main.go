package main

import "fmt"

func main() {
	fmt.Println("hello")
	//var card string = "Cards"
	//card := "Cards"
	//card := []string{"one of spades", "one of hearts", NewDeck()}
	//card := deck{"one of spades", "one of hearts"}
	card := NewDeck()

	//fmt.Println(card)
	for _, values := range card {
		fmt.Println(values)

	}
}

// func NewDeck() string {
// 	return "one of diamonds"
// }
