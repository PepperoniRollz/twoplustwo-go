package simulator

import (
	"fmt"
	"testing"

	card "github.com/pepperonirollz/twoplustwo-go/card"
)

func TestSimulator(t *testing.T) {
	deck := card.NewDeck()
	deck.Shuffle()
	fmt.Println(deck.CurrentState)

	// combin provides several ways to work with the combinations of
	// different objects. Combinations generates them directly.
	fmt.Println("Generate list:")
	// n := 48
	// k := 5
	// list := Combinations(n, k)
	// This is easy, but the number of combinations  can be very large,
	// and generating all at once can use a lot of memory.

	player1 := card.NewCardSet(0)
	player2 := card.NewCardSet(0)
	player1.FromString("AhAd")
	player2.FromString("KsKc")
	deck.RemoveCard(player1.Get(0))
	deck.RemoveCard(player1.Get(1))
	deck.RemoveCard(player2.Get(0))
	deck.RemoveCard(player2.Get(1))
	// fmt.Println(deck.CurrentState)

	// combinations := GenerateCombos(deck.CurrentState, 5)
	// fmt.Println(len(combinations))
	// for i, v := range combinations {
	// 	fmt.Println(i, v)
	// }
	PreFlopEquity([]card.CardSet{player1, player2}, deck.CurrentState)

	// Print the combinations
	// for i, combo := range combinations {
	// 	fmt.Printf("Combination %d: %v\n", i+1, combo)
	// }

	// //make 2 players and deal them both 2 cards
	// players := []d.Player{d.NewPlayer(), d.NewPlayer()}
	// players[0].Hand = []d.Card{deck.DealOne(), deck.DealOne()}
	// players[1].Hand = []d.Card{deck.DealOne(), deck.DealOne()}
	// players[0].Hand = []d.Card{deck.DealOne(), deck.DealOne()}
	// players[1].Hand = []d.Card{deck.DealOne(), deck.DealOne()}
	// // d.PrintHand(players[0].Hand)
	// // d.PrintHand(players[1].Hand)

	// simulator := NewSimulator(deck, players)
	// startTime := time.Now()
	// numSims := 1000000
	// simulator.Simulate(numSims, 5)
	// elapsedTime := time.Since(startTime)
	// fmt.Printf("\nSimulated %d handsin %v!\n", numSims, elapsedTime)
}
