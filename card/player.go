package card

type Player struct {
	Hand   []Card
	Equity float64
}

func NewPlayer() Player {
	return Player{Hand: []Card{}}
}

func PrintHand(hand []Card) {
	for i := 0; i < len(hand); i++ {
		hand[i].Print()
	}
}
