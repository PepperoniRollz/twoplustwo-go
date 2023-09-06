package simulator

import (
	"fmt"
	"testing"
	"time"

	d "github.com/pepperonirollz/twoplustwo-go/deck"
)

func TestSimulator(t *testing.T) {
	deck := d.NewDeck()
	d.ShuffleDeck(deck.InitState)
	//make 2 players and deal them both 2 cards
	players := []d.Player{d.NewPlayer(), d.NewPlayer()}
	players[0].Hand = []d.Card{deck.Deal(), deck.Deal()}
	players[1].Hand = []d.Card{deck.Deal(), deck.Deal()}
	players[0].Hand = []d.Card{deck.Deal(), deck.Deal()}
	players[1].Hand = []d.Card{deck.Deal(), deck.Deal()}
	// d.PrintHand(players[0].Hand)
	// d.PrintHand(players[1].Hand)

	simulator := NewSimulator(deck, players)
	startTime := time.Now()
	numSims := 1000000
	simulator.Simulate(numSims, 5)
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nSimulated %d handsin %v!\n", numSims, elapsedTime)
}
