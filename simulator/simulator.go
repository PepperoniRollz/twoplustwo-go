package simulator

import (
	"fmt"

	card "github.com/pepperonirollz/twoplustwo-go/card"
	combos "github.com/pepperonirollz/twoplustwo-go/combinations"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

// EquityEvaluator takes a slice of hole cards and a board and returns an EquityEvaluation
func EquityEvaluator(holeCards []card.CardSet, board card.CardSet) {
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
	evaluator := eval.NewEvaluator("../HandRanks.dat")

	for i := 0; i < len(combinations); i++ {
		handEvals := make([]eval.HandEvaluation, len(holeCards))

		for j := 0; j < len(holeCards); j++ {
			hand := holeCards[j]
			hand.AddCards(combinations[i])
			handEvals[j] = evaluator.GetHandValue(hand)
		}
		equityEval.EvaluateEquities(handEvals)
		// j := 0
		// if evaluator.CompareHands(handEvals[j], handEvals[j+1]) > 0 {
		// 	wins[j]++
		// } else if evaluator.CompareHands(handEvals[j], handEvals[j+1]) < 0 {
		// 	wins[j+1]++
		// } else {
		// 	ties++
		// }
	}

	equityEval.CalculateEquities()
	equityEval.PrintEquities()
	// for i := 0; i < len(holeCards); i++ {
	// 	equities[i] = float64(wins[i]) / float64(len(combinations))
	// }
	// fmt.Println(equities)
	fmt.Println(len(combinations))
	// fmt.Println(float64(ties) / float64(len(combinations)))
}
