package twoplustwogo

import "fmt"

type Card struct {
	Value int
}

func FromCode(v int) Card {
	return Card{Value: v}
}

func NewCard(s string) Card {
	return parseString(s)
}

func parseString(s string) Card {

	val, ok := GetCardCode()[s]
	if ok {
		return Card{Value: val}
	} else {
		panic("Invalid card string")
	}
}

func (card *Card) Print() {
	fmt.Print(GetCardMap()[card.Value])
}

func (card *Card) ToString() string {
	return GetCardMap()[card.Value]
}
