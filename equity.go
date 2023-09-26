package twoplustwogo

var evaluator Evaluator

func init() {
	evaluator = NewEvaluator("./HandRanks.dat")
}

// EquityEvaluator takes a slice of hole cards and a board and returns an EquityEvaluation
func EvaluateEquity(holeCards []CardSet, board CardSet) EquityEvaluation {
	deck := NewDeck()
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
	numCardsInRunout := 5 - board.Length()
	combinations := GenerateCombos(deckCardSet, numCardsInRunout)

	for i := 0; i < len(combinations); i++ {
		handEvals := make([]HandEvaluation, len(holeCards))

		for j := 0; j < len(holeCards); j++ {
			hand := holeCards[j]
			hand.AddCards(combinations[i])
			handEvals[j] = Evaluate(hand)
		}
		equityEval.EvaluateEquities(handEvals)
	}

	equityEval.CalculateEquities()
	return equityEval
}
