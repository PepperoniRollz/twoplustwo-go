package deck

import (
	c "github.com/pepperonirollz/twoplustwo-go/constants"
)

type Card struct {
	Value int
}

func NewCard(value int) Card {
	return Card{Value: value}
}

func (card *Card) Print() {
	println(c.CARD_MAP[card.Value])
}
