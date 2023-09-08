package twoplustwogo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluator(t *testing.T) {
	evaluator := NewEvaluator("../HandRanks.dat")
	fmt.Println("Initialization complete.")
	h1 := NewHand("TsJsQsKsAs")
	h2 := NewHand("9sTsJsQsKs")

	result1 := evaluator.GetHandValue(h1)
	result2 := evaluator.GetHandValue(h2)
	assert.Greater(t, result1.Value, result2.Value, "h1 should be greater than h2")

	fmt.Println(evaluator.CompareHands(result1, result2))
	result1.Print()
	result2.Print()

}
