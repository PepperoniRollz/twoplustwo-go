package twoplustwogo

import (
	"math/rand"
)

// deck is primarily a wrapper around  cardset.  Also stores initial and current states for monte carlo sims
type Deck struct {
	InitState    CardSet
	CurrentState CardSet
}

func NewDeck() Deck {
	var deck Deck
	for i := 1; i < 53; i++ {
		deck.InitState.AddCard(FromCode(i))
	}
	deck.CurrentState = deck.InitState
	return deck
}

func (d *Deck) Shuffle(r *rand.Rand) {
	for i := 0; i < d.CurrentState.Length(); i++ {
		j := rand.Intn(i + 1)
		d.CurrentState.swap(i, j)
	}
}

func (d *Deck) Deal(n int) CardSet {
	var cards CardSet
	for i := 0; i < n; i++ {
		cards.AddCard(d.CurrentState.Get(i))
	}
	d.CurrentState.Cards = d.CurrentState.Cards[n:]
	return cards
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
