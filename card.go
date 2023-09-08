package twoplustwogo

type Card struct {
	Value       int
	ValueString string
}

func FromCode(v int) Card {
	return Card{Value: v}
}

func FromString(s string) Card {

	val, ok := GetCardCode()[s]
	if ok {
		return Card{Value: val}
	} else {
		panic("Invalid card string")
	}
}

func (card *Card) Print() {
	println(GetCardMap()[card.Value])
}

func (card *Card) ToString() string {
	return GetCardMap()[card.Value]
}
