package card

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

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
func (d deck) Print() {
	for _, values := range d {
		fmt.Println(values)

	}

}
func (d deck) SaveToFile(fileName string) {
	s := strings.Join(d, ",")
	ioutil.WriteFile(fileName, []byte(s), 0644)
}

func (d deck) PickOne() string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(d))
	selectedCard := d[i]
	return selectedCard

}
