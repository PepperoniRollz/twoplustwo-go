package simulator

import (
	"math/rand"
	"testing"
	"time"

	card "github.com/pepperonirollz/twoplustwo-go/card"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func TestSimulator(t *testing.T) {
	players := make([]card.CardSet, r.Intn(10)+2)
	deck := card.NewDeck()
	deck.Shuffle()
	for i := 0; i < len(players); i++ {
		players[i] = card.FromCards([]card.Card{deck.DealOne(), deck.DealOne()})
	}

	board := card.EmptyCardSet()

	equityEval := EquityEvaluator(players, board)
	equityEval.Print()
}

//a function that takes in a slice of HandEvaluations and returns a slice of EquityEvaluations
