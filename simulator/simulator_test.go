package simulator

import (
	"testing"

	card "github.com/pepperonirollz/twoplustwo-go/card"
)

func TestSimulator(t *testing.T) {

	player1 := card.NewHand("AhAd")
	player2 := card.NewHand("KsKc")
	board := card.EmptyCardSet()

	EquityEvaluator([]card.CardSet{player1, player2}, board)
}
