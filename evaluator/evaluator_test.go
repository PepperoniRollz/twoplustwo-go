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
	h1 := card.NewHand("2c2d2h2sAd")
	h2 := card.NewHand("3c3d3h3sAd")

	result1 := evaluator.GetHandValue(h1)
	result2 := evaluator.GetHandValue(h2)
	assert.Equal(t, -1, evaluator.CompareHands(result1, result2), "correct comparison result") // 2222A < 3333A

	fmt.Println(evaluator.CompareHands(result1, result2))
	result1.Print()
	result2.Print()

}
