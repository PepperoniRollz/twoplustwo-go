package twoplustwogo

import (
	"fmt"
)

type HandEvaluation struct {
	HandCategory       int
	RankWithinCategory int
	Hand               string
	Value              int64
}

func newHandEval(p int64, pCards CardSet) HandEvaluation {
	return HandEvaluation{
		HandCategory:       int(p >> 12),
		RankWithinCategory: int(p & 0x00000FFF),
		Value:              p,
	}
}

func (h *HandEvaluation) Print() {
	fmt.Println(h.Hand, GetHandTypes()[h.HandCategory], h.RankWithinCategory, h.Value)
}

func Best5(cards CardSet) CardSet {
	var best CardSet
	var bestHandEval int64 = -1
	combos := GenerateCombos(cards, 5)
	for i := 0; i < len(combos); i++ {
		handValue := evaluator.GetHandValue(combos[i])
		if handValue.Value > bestHandEval {
			best = combos[i]
		}
	}
	return best
}
