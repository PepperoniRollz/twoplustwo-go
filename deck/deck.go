package deck

import (
	"math/rand"
)

type Deck struct {
	InitState    []Card
	CurrentState []Card
}

func NewDeck() Deck {
	var deck Deck
	deck.InitState = make([]Card, 53)
	for i := 1; i < 53; i++ {
		deck.InitState[i] = NewCard(i)
	}
	deck.CurrentState = deck.InitState
	return deck
}

func ShuffleDeck(deck []Card) {
	for i := 0; i < len(deck); i++ {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

// func (d *Deck) Shuffle() {
// 	for i := 1; i < 53; i++ {
// 		j := rand.Intn(i + 1)
// 		d.CurrentState[i], d.CurrentState[j] = d.CurrentState[j], d.CurrentState[i]
// 	}
// }

func (d *Deck) Deal() Card {
	var card Card = d.CurrentState[0]
	d.CurrentState = d.CurrentState[1:]
	return card
}
