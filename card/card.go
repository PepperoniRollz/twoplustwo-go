package card

import (
	c "github.com/pepperonirollz/twoplustwo-go/constants"
)

type Card struct {
	Value       int
	ValueString string
}

func FromCode(v int) Card {
	return Card{Value: v}
}

func FromString(s string) Card {

	val, ok := c.CARD_CODE[s]
	if ok {
		return Card{Value: val}
	} else {
		panic("Invalid card string")
	}
}

func (card *Card) Print() {
	println(c.CARD_MAP[card.Value])
}
