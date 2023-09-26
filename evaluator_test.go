package twoplustwogo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var gopoker5 = []CardSet{
	NewHand("AsAcJc7h5d"), // pair
	NewHand("AsAcJcJd5d"), // two pair
	NewHand("AsAcAdJd5d"), // three of a kind
	NewHand("AsKsQdJhTd"), // straight
	NewHand("Ts7s4s3s2s"), // flush
	NewHand("4s4c4d2s2h"), // full house
	NewHand("AsAcAdAh5h"), // four of a kind
	NewHand("AsKsQsJsTs"), // straight flush
}

var gopoker6 = []CardSet{
	NewHand("3dAsKsJc7h5d"), // high card
	NewHand("3dAsAcJc7h5d"), // pair
	NewHand("3dAsAcJcJd5d"), // two pair
	NewHand("3dAsAcAdJd5d"), // three of a kind
	NewHand("3dAsKsQdJhTd"), // straight
	NewHand("3dTs7s4s3s2s"), // flush
	NewHand("3d4s4c4d2s2h"), // full house
	NewHand("3dAsAcAdAh5h"), // four of a kind
	NewHand("3dAsKsQsJsTs"), // straight flush
}

var gopoker7 = []CardSet{
	NewHand("3dAsKsJc7h5d2d"), // high card
	NewHand("3dAsAcJc7h5d2d"), // pair
	NewHand("3dAsAcJcJd5d2d"), // two pair
	NewHand("3dAsAcAdJd5d2d"), // three of a kind
	NewHand("3dAsKsQdJhTd2d"), // straight
	NewHand("3dTs7s4s3s2s2d"), // flush
	NewHand("3d4s4c4d2s2h2d"), // full house
	NewHand("3dAsAcAdAh5h2d"), // four of a kind
	NewHand("3dAsKsQsJsTs2d"), // straight flush
}

var deck = NewDeck()

var hands []CardSet = GenerateCombos(deck.CurrentState, 5)

func TestEvaluator(t *testing.T) {
	fmt.Println("Initialization complete.")
	h1 := gopoker6[0]
	h2 := gopoker6[1]
	h3 := gopoker7[6]
	h4 := gopoker6[7]

	result1 := Evaluate(h1)
	result2 := Evaluate(h2)
	result3 := Evaluate(h3)
	result4 := Evaluate(h4)
	result5 := Evaluate5(h3)
	result6 := Evaluate5(h4)

	assert.Less(t, result1.Value, result2.Value, "h1 should be worse than h2")
	assert.Less(t, result3.Value, result4.Value, "h3 should be worse than h4")

	result3.Print()
	result4.Print()
	result5.Print()
	result6.Print()

}

func benchmarkEvaluate5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, hand := range gopoker5 {
			Evaluate(hand)
		}
	}
}

func benchmarkEvaluate6(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, hand := range gopoker6 {
			Evaluate(hand)
		}
	}
}

func benchmarkEvaluate7(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, hand := range gopoker7 {
			Evaluate(hand)
		}
	}
}

func benchmarkAll5CardHands(b *testing.B) {
	//split the hands into two equal parts
	// this generates all C(52,5) = 2,598,960 possible hands
	hands1 := hands[:len(hands)/2]
	hands2 := hands[len(hands)/2:]

	//compares all the hands against each other
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(hands1); i++ {
			hand1eval := Evaluate(hands1[i])
			hand2eval := Evaluate(hands2[i])
			CompareHands(hand1eval, hand2eval)
		}
	}

}

func benchmarkAll6CardHands(b *testing.B) {
	//split the hands into two equal parts
	// this generates all C(52,6) = 20,358,520 possible hands
	hands1 := hands[:len(hands)/2]
	hands2 := hands[len(hands)/2:]

	//compares all the hands against each other
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(hands1); i++ {
			hand1eval := Evaluate(hands1[i])
			hand2eval := Evaluate(hands2[i])
			CompareHands(hand1eval, hand2eval)
		}
	}

}

func benchmarkAll7CardHands(b *testing.B) {
	//split the hands into two equal parts
	// this generates all C(52,7) = 133,784,560 possible hands
	hands1 := hands[:len(hands)/2]
	hands2 := hands[len(hands)/2:]

	//compares all the hands against each other

	for n := 0; n < b.N; n++ {
		for i := 0; i < len(hands1); i++ {
			hand1eval := Evaluate(hands1[i])
			hand2eval := Evaluate(hands2[i])
			CompareHands(hand1eval, hand2eval)
		}
	}
}

func BenchmarkGoPoker5(b *testing.B) { benchmarkEvaluate5(b) }
func BenchmarkGoPoker6(b *testing.B) { benchmarkEvaluate6(b) }
func BenchmarkGoPoker7(b *testing.B) { benchmarkEvaluate7(b) }

func BenchmarkAll5CardHands(b *testing.B) { benchmarkAll5CardHands(b) }
func BenchmarkAll6CardHands(b *testing.B) { benchmarkAll6CardHands(b) }
func BenchmarkAll7CardHands(b *testing.B) { benchmarkAll7CardHands(b) }
