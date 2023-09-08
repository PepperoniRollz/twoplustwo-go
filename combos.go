package twoplustwogo

func generate(deck, current CardSet, k, index int, combinations *[]CardSet) {
	if current.Length() == k {
		set := EmptyCardSet()
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
func GenerateCombos(deck CardSet, k int) []CardSet {
	var combinations []CardSet
	generate(deck, CardSet{}, k, 0, &combinations)
	return combinations
}
