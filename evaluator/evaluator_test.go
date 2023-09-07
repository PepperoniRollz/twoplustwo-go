package evaluator

import (
	"fmt"
	"testing"

	card "github.com/pepperonirollz/twoplustwo-go/card"
)

func TestEvaluator(t *testing.T) {
	evaluator := NewEvaluator("../HandRanks.dat")
	fmt.Println("Initialization complete.")
	h1 := card.NewHand("TsJsQsKsAs")
	h2 := card.NewHand("9sTsJsQsKs")

	result1 := evaluator.GetHandValue(h1)
	result2 := evaluator.GetHandValue(h2)
	// assert.Equal(t, -1, evaluator.CompareHands(result1, result2), "correct comparison result") // 2222A < 3333A

	fmt.Println(evaluator.CompareHands(result1, result2))
	result1.Print()
	result2.Print()

}
