package equityEval

import (
	"fmt"

	"github.com/pepperonirollz/twoplustwo-go/card"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

type EquityEvaluation struct {
	Equities  []float64
	TieEquity []float64
	Wins      []int
	Losses    []int
	Ties      []int
	HoleCards []card.CardSet
	Board     card.CardSet
}

func NewEquityEvaluation(holeCards []card.CardSet, board card.CardSet) EquityEvaluation {
	size := len(holeCards)
	hc := make([]card.CardSet, len(holeCards))
	copy(hc, holeCards)
	return EquityEvaluation{
		Equities:  make([]float64, size),
		TieEquity: make([]float64, size),
		Ties:      make([]int, size),
		Wins:      make([]int, size),
		Losses:    make([]int, size),
		HoleCards: hc,
		Board:     board,
	}
}

func (e *EquityEvaluation) EvaluateEquities(handEvals []eval.HandEvaluation) {
	size := len(handEvals)
	var maxScore int64 = 0
	tieCount := 0
	//get max score
	for i := 0; i < size; i++ {
		if handEvals[i].Value > maxScore {
			maxScore = handEvals[i].Value
		}
	}
	//see if there are any ties for best hand
	for i := 0; i < size; i++ {
		if handEvals[i].Value == maxScore {
			tieCount++
		}
	}
	//if there are ties, then the best hands will get a tie, all others get losses
	if tieCount > 1 {
		for i := 0; i < size; i++ {
			if handEvals[i].Value == maxScore {
				e.Ties[i]++
			} else {
				e.Losses[i]++
			}
		}
		//if there are no ties, then the best hand gets a win, all others get losses
	} else {
		for i := 0; i < size; i++ {
			if handEvals[i].Value == maxScore {
				e.Wins[i]++
			} else {
				e.Losses[i]++
			}
		}
	}
}

//this does not account for the instance where, for example, one hand wins, and the other 2 hands are ties/chops.  In this scenario, there is one winner and 2 losers.

func (e *EquityEvaluation) CalculateEquities() {
	size := len(e.HoleCards)
	for i := 0; i < size; i++ {
		e.Equities[i] = float64(e.Wins[i]) / float64(e.Wins[i]+e.Losses[i]+e.Ties[i])
		e.TieEquity[i] = float64(e.Ties[i]) / float64(e.Wins[i]+e.Losses[i]+e.Ties[i])
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
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf(" %-10s | %-10s |%-10s | %-10s | %-10s | %-10s | %-10s |\n", "Hole Cards", "Board", "Equity", "TieEquity", "Wins", "Losses", "Ties")
	fmt.Println("----------------------------------------------------------------------------------")

	for i := range e.Equities {
		holeCardsStr := ""
		if len(e.HoleCards) > i {
			holeCardsStr = e.HoleCards[i].ToString()
		}

		fmt.Printf(" %-10s | %-10s |%-10.4f | %-10.4f | %-10d | %-10d | %-10d | \n", holeCardsStr, e.Board.ToString(), e.Equities[i], e.TieEquity[i], e.Wins[i], e.Losses[i], e.Ties[i])
	}

	fmt.Println("-----------------------------------------------------------------------------------")
}
