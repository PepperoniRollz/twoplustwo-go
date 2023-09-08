package twoplustwogo

import (
	"testing"
)

var deck = NewDeck()

func BenchmarkEvaluator(b *testing.B) {
	//split the hands into two equal parts
	// this generates all C(52,5) = 2,598,960 possible hands
	var hands []CardSet = GenerateCombos(deck.CurrentState, 5)
	hands1 := hands[:len(hands)/2]
	hands2 := hands[len(hands)/2:]

	//compares all the hands against each other
	for i := 0; i < len(hands1); i++ {
		hand1eval := evaluator.GetHandValue(hands1[i])
		hand2eval := evaluator.GetHandValue(hands2[i])

		evaluator.CompareHands(hand1eval, hand2eval)
	}

}
