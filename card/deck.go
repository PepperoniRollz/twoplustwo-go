package card

import (
	"math/rand"
)

type Deck struct {
	InitState    CardSet
	CurrentState CardSet
}

func NewDeck() Deck {
	var deck Deck
	deck.InitState = NewCardSet(52)
	for i := 1; i < 53; i++ {
		deck.InitState.Set(i-1, FromCode(i))
	}
	deck.CurrentState = deck.InitState
	return deck
}

func (d *Deck) Shuffle() {
	for i := 0; i < d.CurrentState.Length(); i++ {
		j := rand.Intn(i + 1)
		d.CurrentState.swap(i, j)
	}
}

func (d *Deck) DealOne() Card {
	var card Card = d.CurrentState.Get(0)
	d.CurrentState.Cards = d.CurrentState.Cards[1:]
	return card
}

func (d *Deck) RemoveCard(card Card) {
	for i := 0; i < d.CurrentState.Length(); i++ {
		if d.CurrentState.Get(i) == card {
			d.CurrentState.Cards = append(d.CurrentState.Cards[:i], d.CurrentState.Cards[i+1:]...)
			return
		}
	}
}

func (d *Deck) RemoveCards(cards CardSet) {
	for i := 0; i < cards.Length(); i++ {
		d.RemoveCard(cards.Get(i))
	}
}
