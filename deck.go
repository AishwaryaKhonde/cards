package main

type deck []string

func NewDeck() deck {
	card := []string{}
	cardsType := []string{"spades", "hearts", "diamonds", "clubs"}
	cardsValues := []string{"one", "two", "three", "four"}

	for _, cardType := range cardsType {
		for _, cardsValues := range cardsValues {
			card = append(card, cardsValues+" of "+cardType)

		}

	}
	return card
}
