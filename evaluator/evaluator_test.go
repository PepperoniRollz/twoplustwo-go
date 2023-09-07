package evaluator

import (
	"fmt"
	"testing"

	card "github.com/pepperonirollz/twoplustwo-go/card"
	"github.com/stretchr/testify/assert"
)

func TestEvaluator(t *testing.T) {
	evaluator := NewEvaluator("../HandRanks.dat")
	fmt.Println("Initialization complete.")
	var q1, q2 card.CardSet
	q1.FromString("2c2d2h2sAd")
	q2.FromString("3c3d3h3sAd")
	// quad2s := []card.Card{card.FromString("2c"), card.FromString("2d"), card.FromString("2h"), card.FromString("2s"), card.FromString("Ad")}
	// quad3s := []card.Card{card.FromString("3c"), card.FromString("3d"), card.FromString("3h"), card.FromString("3s"), card.FromString("Ad")}
	// straight7WithK := []card.Card{card.FromString("3h"), card.FromString("4h"), card.FromString("5h"), card.FromString("6h"), card.FromString("7d"), card.FromString("Kd")}
	// straight7WithA := []card.Card{card.FromString("3h"), card.FromString("4h"), card.FromString("5h"), card.FromString("6h"), card.FromString("7d"), card.FromString("Ad")}
	// strightFlush5 := []d.Card{d.NewCard(1), d.NewCard(5), d.NewCard(9), d.NewCard(4), d.NewCard(26)}

	result1 := evaluator.GetHandValue(q1)
	result2 := evaluator.GetHandValue(q2)
	// rseult3 := evaluator.GetHandValue(straight7WithK)
	// result4 := evaluator.GetHandValue(straight7WithA)
	assert.Equal(t, -1, evaluator.CompareHands(result1, result2), "correct comparison result") // 2222A < 3333A
	// assert.Equal(t, 0, evaluator.CompareHands(rseult3, result4), "correct comparison result")  //345673K == 345673A

	fmt.Println(evaluator.CompareHands(result1, result2))
	result1.Print()
	result2.Print()

}
