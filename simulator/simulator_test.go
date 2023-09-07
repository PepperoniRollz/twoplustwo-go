package simulator

import (
	"testing"

	card "github.com/pepperonirollz/twoplustwo-go/card"
)

func TestSimulator(t *testing.T) {

	player1 := card.NewHand("AhAd")
	player2 := card.NewHand("KsKc")
	player3 := card.NewHand("QsQc")
	player4 := card.NewHand("JsJc")
	// board := card.NewBoard("5h6h7s")
	board := card.EmptyCardSet()

	equityEval := EquityEvaluator([]card.CardSet{player1, player2, player3, player4}, board)
	equityEval.Print()
}

//a function that takes in a slice of HandEvaluations and returns a slice of EquityEvaluations
