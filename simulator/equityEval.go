package simulator

import (
	"fmt"

	"github.com/pepperonirollz/twoplustwo-go/card"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

type EquityEvaluation struct {
	Equities  []float64
	Ties      []int
	Wins      []int
	Losses    []int
	HoleCards []card.CardSet
	Board     card.CardSet
}

func NewEquityEvaluation(holeCards []card.CardSet, board card.CardSet) EquityEvaluation {
	size := len(holeCards)
	return EquityEvaluation{
		Equities:  make([]float64, size),
		Ties:      make([]int, size),
		Wins:      make([]int, size),
		Losses:    make([]int, size),
		HoleCards: holeCards,
		Board:     board,
	}
}

func (e *EquityEvaluation) EvaluateEquities(handEvals []eval.HandEvaluation) {
	size := len(handEvals)
	var maxScore int64 = 0
	tieCount := 0
	//get max score
	for i := 0; i < size; i++ {
		if handEvals[i].P > maxScore {
			maxScore = handEvals[i].P
		}
	}
	//see if there are any ties for best hand
	for i := 0; i < size; i++ {
		if handEvals[i].P == maxScore {
			tieCount++
		}
	}
	//if there are ties, then the best hands will get a tie, all others get losses
	if tieCount > 1 {
		for i := 0; i < size; i++ {
			if handEvals[i].P == maxScore {
				e.Ties[i]++
			} else {
				e.Losses[i]++
			}
		}
		//if there are no ties, then the best hand gets a win, all others get losses
	} else {
		for i := 0; i < size; i++ {
			if handEvals[i].P == maxScore {
				e.Wins[i]++
			} else {
				e.Losses[i]++
			}
		}
	}
}

func (e *EquityEvaluation) CalculateEquities() {
	size := len(e.HoleCards)
	for i := 0; i < size; i++ {
		e.Equities[i] = float64(e.Wins[i]) / float64(e.Wins[i]+e.Losses[i]+e.Ties[i])
	}
}

func (e *EquityEvaluation) PrintEquities() {
	size := len(e.HoleCards)
	for i := 0; i < size; i++ {
		e.HoleCards[i].Print()
	}
	fmt.Println(e)
}

func (e *EquityEvaluation) Print() {
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("| %-10s | %-10s | %-10s | %-10s | %-20s |\n", "Equity", "Wins", "Losses", "Ties", "Hole Cards")
	fmt.Println("--------------------------------------------------------")

	for i := range e.Equities {
		holeCardsStr := ""
		if len(e.HoleCards) > i {
			holeCardsStr = e.HoleCards[i].ToString()
		}

		fmt.Printf("| %-10.2f | %-10d | %-10d | %-10d | %-20s |\n", e.Equities[i], e.Wins[i], e.Losses[i], e.Ties[i], holeCardsStr)
	}

	fmt.Println("--------------------------------------------------------")
}
