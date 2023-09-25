package twoplustwogo

import (
	"fmt"
)

type HandEvaluation struct {
	HandCategory       int
	RankWithinCategory int
	Hand               CardSet
	Value              int64
}

func newHandEval(p int64, pCards CardSet) HandEvaluation {
	return HandEvaluation{
		HandCategory:       int(p >> 12),
		RankWithinCategory: int(p & 0x00000FFF),
		Value:              p,
		Hand:               pCards,
	}
}

func (h *HandEvaluation) Print() {
	fmt.Println(GetHandTypes()[h.HandCategory], h.RankWithinCategory, h.Value)
	h.Hand.Print()
}
