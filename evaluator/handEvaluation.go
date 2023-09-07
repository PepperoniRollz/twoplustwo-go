package evaluator

import (
	"fmt"

	d "github.com/pepperonirollz/twoplustwo-go/card"
	c "github.com/pepperonirollz/twoplustwo-go/constants"
)

type HandEvaluation struct {
	HandCategory       int
	RankWithinCategory int
	Hand               string
	P                  int64
}

func newHandEval(p int64, pCards []d.Card) HandEvaluation {

	return HandEvaluation{
		HandCategory:       int(p >> 12),
		RankWithinCategory: int(p & 0x00000FFF),
		Hand:               cardsToString(pCards),
		P:                  p,
	}
}

func (h *HandEvaluation) Print() {
	fmt.Println(h.Hand, c.HAND_TYPES[h.HandCategory], h.RankWithinCategory, h.P)

}

func cardsToString(cards []d.Card) string {
	var s string
	for i := 0; i < len(cards); i++ {
		s += c.CARD_MAP[cards[i].Value] + " "
	}
	return s
}
