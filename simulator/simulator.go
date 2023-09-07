package simulator

import (
	"fmt"

	card "github.com/pepperonirollz/twoplustwo-go/card"
	eval "github.com/pepperonirollz/twoplustwo-go/evaluator"
)

type Simulator struct {
	Deck    card.Deck
	Players []card.Player
}

func NewSimulator(deck card.Deck, players []card.Player) Simulator {
	return Simulator{Deck: deck, Players: players}
}

// func (s *Simulator) Simulate(numSims int, cardsLeft int) {
// 	evaluator := eval.NewEvaluator("../HandRanks.dat")
// 	p1Wins := 0
// 	p2Wins := 0
// 	ties := 0
// 	card.PrintHand(s.Players[0].Hand)
// 	card.PrintHand(s.Players[1].Hand)
// 	for i := 0; i < numSims; i++ {
// 		tmpHand1 := s.Players[0].Hand
// 		tmpHand2 := s.Players[1].Hand

// 		for k := 0; k < cardsLeft; k++ {
// 			newCard := s.Deck.DealOne()
// 			// fmt.Println("New card:", newCard)
// 			// for j := 0; j < len(s.Players); j++ {
// 			// 	s.Players[j].Hand = append(s.Players[j].Hand, newCard)
// 			// }
// 			tmpHand1 = append(tmpHand1, newCard)
// 			tmpHand2 = append(tmpHand2, newCard)
// 		}
// 		result1 := evaluator.GetHandValue(tmpHand1)
// 		result2 := evaluator.GetHandValue(tmpHand2)
// 		if evaluator.CompareHands(result1, result2) > 0 {
// 			p1Wins++
// 		} else if evaluator.CompareHands(result1, result2) < 0 {
// 			p2Wins++
// 		} else {
// 			ties++
// 		}
// 		// result1.Print()
// 		// result2.Print()

// 		// fmt.Printf("\rSims - %d", i)
// 		fmt.Println(s.Deck.CurrentState)
// 		// card.Shuffle()
// 		// card.ShuffleDeck(s.Deck.CurrentState)
// 	}
// 	println("Player 1 wins:", p1Wins, "Player 2 wins:", p2Wins, "Ties:", ties)
// 	fmt.Printf("%f wins\n", float64(p1Wins)/float64(numSims))
// 	fmt.Printf("%f losses\n", float64(p2Wins)/float64(numSims))
// 	fmt.Printf("%f ties\n", float64(ties)/float64(numSims))
// }

func generate(deck, current card.CardSet, k, index int, combinations *[]card.CardSet) {
	if current.Length() == k {
		set := card.NewCardSet(0)
		set.AddCards(current)
		*combinations = append(*combinations, set)

		// *combinations = append(*combinations current)
		return
	}
	if index >= deck.Length() {
		return
	}

	current.AddCard(deck.Get(index))
	generate(deck, current, k, index+1, combinations)
	current.Cards = current.Cards[:current.Length()-1]
	generate(deck, current, k, index+1, combinations)
}

func GenerateCombos(deck card.CardSet, k int) []card.CardSet {
	var combinations []card.CardSet
	generate(deck, card.CardSet{}, k, 0, &combinations)
	return combinations
}

func Combinations(n, k int) [][]int {
	combins := Binomial(n, k)
	data := make([][]int, combins)
	if len(data) == 0 {
		return data
	}
	data[0] = make([]int, k)
	for i := range data[0] {
		data[0][i] = i
	}
	for i := 1; i < combins; i++ {
		next := make([]int, k)
		copy(next, data[i-1])
		nextCombination(next, n, k)
		data[i] = next
	}
	return data
}

func Binomial(n, k int) int {
	if n < 0 || k < 0 {
		panic("errNegInput")
	}
	if n < k {
		panic("badSetSize")
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

func nextCombination(s []int, n, k int) {
	for j := k - 1; j >= 0; j-- {
		if s[j] == n+j-k {
			continue
		}
		s[j]++
		for l := j + 1; l < k; l++ {
			s[l] = s[j] + l - j
		}
		break
	}
}

func PreFlopEquity(holeCards []card.CardSet, deck card.CardSet) {
	wins := make([]int, len(holeCards))
	ties := 0
	equities := make([]float64, len(holeCards))
	combinations := GenerateCombos(deck, 5)
	evaluator := eval.NewEvaluator("../HandRanks.dat")

	for i := 0; i < len(combinations); i++ {
		handEvals := make([]eval.HandEvaluation, len(holeCards))
		// fmt.Println("combo:", i, " ", combinations[i])

		for j := 0; j < len(holeCards); j++ {
			hand := holeCards[j]
			hand.AddCards(combinations[i])
			handEvals[j] = evaluator.GetHandValue(hand)
			fmt.Println("j: ", j, " combo", i, " hand:", holeCards[j], "combo:", combinations[i], "eval:", handEvals[j])
		}
		j := 0
		if evaluator.CompareHands(handEvals[j], handEvals[j+1]) > 0 {
			wins[j]++
		} else if evaluator.CompareHands(handEvals[j], handEvals[j+1]) < 0 {
			wins[j+1]++
		} else {
			ties++
		}
	}
	for i := 0; i < len(holeCards); i++ {
		equities[i] = float64(wins[i]) / float64(len(combinations))
	}
	fmt.Println(equities)
	fmt.Println(len(combinations))
	fmt.Println(float64(ties) / float64(len(combinations)))
}

//use --> generateCombinations(deck, 5) gets all 5 card combos from the remaining deck, should pass in remaing deck after initial hands are dealt
