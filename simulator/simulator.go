package simulator

import (
	"fmt"

	d "github.com/pepperonirollz/twoplustwo-go/deck"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

type Simulator struct {
	Deck    d.Deck
	Players []d.Player
}

func NewSimulator(deck d.Deck, players []d.Player) Simulator {
	return Simulator{Deck: deck, Players: players}
}

func (s *Simulator) Simulate(numSimlations int, cardsLeft int) {
	evaluator := eval.NewEvaluator()
	p1Wins := 0
	p2Wins := 0
	ties := 0
	d.PrintHand(s.Players[0].Hand)
	d.PrintHand(s.Players[1].Hand)
	for i := 0; i < numSimlations; i++ {
		tmpHand1 := s.Players[0].Hand
		tmpHand2 := s.Players[1].Hand

		for k := 0; k < cardsLeft; k++ {
			newCard := s.Deck.Deal()
			// fmt.Println("New card:", newCard)
			// for j := 0; j < len(s.Players); j++ {
			// 	s.Players[j].Hand = append(s.Players[j].Hand, newCard)
			// }
			tmpHand1 = append(tmpHand1, newCard)
			tmpHand2 = append(tmpHand2, newCard)
		}
		result1 := evaluator.GetHandValue(tmpHand1)
		result2 := evaluator.GetHandValue(tmpHand2)
		if evaluator.CompareHands(result1, result2) > 0 {
			p1Wins++
		} else if evaluator.CompareHands(result1, result2) < 0 {
			p2Wins++
		} else {
			ties++
		}
		// result1.Print()
		// result2.Print()

		// fmt.Printf("\rSims - %d", i)
		fmt.Println(s.Deck.CurrentState)
		d.ShuffleDeck(s.Deck.CurrentState)
	}
	println("Player 1 wins:", p1Wins, "Player 2 wins:", p2Wins, "Ties:", ties)
	fmt.Printf("%f wins\n", float64(p1Wins)/float64(numSimlations))
	fmt.Printf("%f losses\n", float64(p2Wins)/float64(numSimlations))
	fmt.Printf("%f ties\n", float64(ties)/float64(numSimlations))
	// println("Player 1 wins", float64(p1Wins)/float64(numSimlations), "Player 2 wins:", float64(p2Wins)/float64(numSimlations), "Ties:", float64(ties)/float64(numSimlations))
}
