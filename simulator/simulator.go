package simulator

import (
	"fmt"

	card "github.com/pepperonirollz/twoplustwo-go/card"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

type Simulator struct {
	Deck    card.Deck
	Players []card.Player
}

func NewSimulator(deck card.Deck, players []card.Player) Simulator {
	return Simulator{Deck: deck, Players: players}
}

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

func GenerateCombos(deck card.CardSet, k int) []card.CardSet {
	var combinations []card.CardSet
	generate(deck, card.CardSet{}, k, 0, &combinations)
	return combinations
}

func EquityEvaluator(holeCards []card.CardSet, board card.CardSet) {

	deck := card.NewDeck()

	for i := 0; i < len(holeCards); i++ {
		deck.RemoveCards(holeCards[i])
	}
	if board.Length() != 0 {
		deck.RemoveCards(board)
		for i := 0; i < len(holeCards); i++ {
			holeCards[i].AddCards(board)
		}
	}
	deckCardSet := deck.CurrentState
	wins := make([]int, len(holeCards))
	ties := 0
	equities := make([]float64, len(holeCards))
	numCardsInRunout := 5 - board.Length()
	combinations := GenerateCombos(deckCardSet, numCardsInRunout)
	evaluator := eval.NewEvaluator("../HandRanks.dat")

	for i := 0; i < len(combinations); i++ {
		handEvals := make([]eval.HandEvaluation, len(holeCards))

		for j := 0; j < len(holeCards); j++ {
			hand := holeCards[j]
			hand.AddCards(combinations[i])
			handEvals[j] = evaluator.GetHandValue(hand)
		}
		j := 0
		if evaluator.CompareHands(handEvals[j], handEvals[j+1]) > 0 {
			wins[j]++
		} else if evaluator.CompareHands(handEvals[j], handEvals[j+1]) < 0 {
			wins[j+1]++
		} else {
			ties++
		}
	}
	for i := 0; i < len(holeCards); i++ {
		equities[i] = float64(wins[i]) / float64(len(combinations))
	}
	fmt.Println(equities)
	fmt.Println(len(combinations))
	fmt.Println(float64(ties) / float64(len(combinations)))
}
