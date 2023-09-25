package twoplustwogo

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Evaluator struct {
	HR []int64
}

func (e *Evaluator) Evaluate(pCards CardSet) HandEvaluation {
	var p int64 = 53
	size := len(pCards.Cards)
	if size < 5 {
		panic("Not enough cards to evaluate hand.")
	}
	if size > 7 {
		panic("Too many cards to evaluate hand.")
	}
	for i := 0; i < size; i++ {
		p = e.HR[p+int64(pCards.Cards[i].Value)]
	}

	if size == 5 || size == 6 {
		p = e.HR[p]
	}

	return newHandEval(p, pCards)
}

func (e *Evaluator) CompareHands(hand1 HandEvaluation, hand2 HandEvaluation) int {

	if hand1.Value > hand2.Value {
		return 1
	} else if hand1.Value < hand2.Value {
		return -1
	} else {
		return 0
	}
}

func NewEvaluator(pathToHandRanks string) Evaluator {
	file, err := os.Open(pathToHandRanks)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	HR := make([]int64, 32487834)
	if err := binary.Read(file, binary.LittleEndian, &HR); err != nil {
		fmt.Println("Error reading HR data:", err)

	}

	return Evaluator{HR: HR}

}

func Best5(cards CardSet) CardSet {
	var best CardSet
	var bestHandEval int64 = -1
	combos := GenerateCombos(cards, 5)
	for i := 0; i < len(combos); i++ {
		handValue := evaluator.Evaluate(combos[i])
		if handValue.Value > bestHandEval {
			best = combos[i]
		}
	}
	return best
}

func (e *Evaluator) Evaluate5(cards CardSet) HandEvaluation {
	fiveBest := Best5(cards)
	return e.Evaluate(fiveBest)
}
