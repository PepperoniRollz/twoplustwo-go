package simulator

import (
	card "github.com/pepperonirollz/twoplustwo-go/card"
	combos "github.com/pepperonirollz/twoplustwo-go/combinations"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

var evaluator eval.Evaluator

func init() {
	evaluator = eval.NewEvaluator("../HandRanks.dat")
}

// EquityEvaluator takes a slice of hole cards and a board and returns an EquityEvaluation
func EquityEvaluator(holeCards []card.CardSet, board card.CardSet) EquityEvaluation {
	deck := card.NewDeck()
	equityEval := NewEquityEvaluation(holeCards, board)

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
	// equities := make([]float64, len(holeCards))
	numCardsInRunout := 5 - board.Length()
	combinations := combos.GenerateCombos(deckCardSet, numCardsInRunout)

	for i := 0; i < len(combinations); i++ {
		handEvals := make([]eval.HandEvaluation, len(holeCards))

		for j := 0; j < len(holeCards); j++ {
			hand := holeCards[j]
			hand.AddCards(combinations[i])
			handEvals[j] = evaluator.GetHandValue(hand)
		}
		equityEval.EvaluateEquities(handEvals)
	}

	equityEval.CalculateEquities()
	return equityEval
}
