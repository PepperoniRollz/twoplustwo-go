package card

import (
	constants "github.com/pepperonirollz/twoplustwo-go/constants"
)

type Card struct {
	Value       int
	ValueString string
}

func FromCode(v int) Card {
	return Card{Value: v}
}

func FromString(s string) Card {

	val, ok := constants.GetCardCode()[s]
	if ok {
		return Card{Value: val}
	} else {
		panic("Invalid card string")
	}
}

func (card *Card) Print() {
	println(constants.GetCardMap()[card.Value])
}

func (card *Card) ToString() string {
	return constants.GetCardMap()[card.Value]
}
