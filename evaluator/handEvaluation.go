package evaluator

import (
	"fmt"

	d "github.com/pepperonirollz/twoplustwo-go/card"
	constants "github.com/pepperonirollz/twoplustwo-go/constants"
)

type HandEvaluation struct {
	HandCategory       int
	RankWithinCategory int
	Hand               string
	Value              int64
}

func newHandEval(p int64, pCards []d.Card) HandEvaluation {

	return HandEvaluation{
		HandCategory:       int(p >> 12),
		RankWithinCategory: int(p & 0x00000FFF),
		Value:              p,
	}
}

func (h *HandEvaluation) Print() {
	fmt.Println(h.Hand, constants.GetHandTypes()[h.HandCategory], h.RankWithinCategory, h.Value)

}

func cardsToString(cards []d.Card) string {
	var s string
	for i := 0; i < len(cards); i++ {
		s += constants.GetCardMap()[cards[i].Value] + " "
	}
	return s
}
