package evaluator

import (
	"fmt"
	"testing"

	d "github.com/pepperonirollz/twoplustwo-go/deck"
)

func TestEvaluator(t *testing.T) {
	evaluator := NewEvaluator()
	fmt.Println("Initialization complete.")

	// h1 := []int{2, 3, 6, 7, 11, 52} // Example card values
	// h2 := []int{2, 3, 6, 7, 10, 52} // Example card values
	var h3 = []d.Card{d.NewCard(2), d.NewCard(3), d.NewCard(6), d.NewCard(7), d.NewCard(48), d.NewCard(52)}
	var h4 = []d.Card{d.NewCard(2), d.NewCard(3), d.NewCard(6), d.NewCard(7), d.NewCard(11), d.NewCard(52)}
	result1 := evaluator.GetHandValue(h3)
	result2 := evaluator.GetHandValue(h4)
	fmt.Println(evaluator.CompareHands(result1, result2))
	result1.Print()
	result2.Print()

}
