package combos

import "github.com/pepperonirollz/twoplustwo-go/card"

func generate(deck, current card.CardSet, k, index int, combinations *[]card.CardSet) {
	if current.Length() == k {
		set := card.EmptyCardSet()
		set.AddCards(current)
		*combinations = append(*combinations, set)

		return
	}
	if index >= deck.Length() {
		return
	}

	current.AddCard(deck.Get(index))
	generate(deck, current, k, index+1, combinations)
	current.Cards = current.Cards[:current.Length()-1]
	generate(deck, current, k, index+1, combinations)
}

// generate C(n,k) combinations of cards from a deck of n cards
// returns a slice of card.CardSet
func GenerateCombos(deck card.CardSet, k int) []card.CardSet {
	var combinations []card.CardSet
	generate(deck, card.CardSet{}, k, 0, &combinations)
	return combinations
}
