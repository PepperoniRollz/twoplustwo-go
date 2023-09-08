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

func newHandEval(p int64, pCards []Card) HandEvaluation {

	return HandEvaluation{
		HandCategory:       int(p >> 12),
		RankWithinCategory: int(p & 0x00000FFF),
		Value:              p,
	}
}

func (h *HandEvaluation) Print() {
	fmt.Println(h.Hand, GetHandTypes()[h.HandCategory], h.RankWithinCategory, h.Value)

}

func cardsToString(cards []Card) string {
	var s string
	for i := 0; i < len(cards); i++ {
		s += GetCardMap()[cards[i].Value] + " "
	}
	return s
}
