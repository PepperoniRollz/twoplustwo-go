package twoplustwogo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkHandsPerSecond(b *testing.B) {
	evaluator := getDefaultEvaluator()
	
	// Generate random 7-card hands for benchmarking
	hands := make([]CardSet, 10000)
	for i := 0; i < 10000; i++ {
		hands[i] = generateRandomHand(7)
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		hand := hands[i%10000]
		_ = evaluator.Evaluate(hand)
	}
}

func TestHandsPerSecondTiming(t *testing.T) {
	evaluator := getDefaultEvaluator()
	numHands := 1000000
	
	// Test 5-card hands
	hands5 := make([]CardSet, numHands)
	for i := 0; i < numHands; i++ {
		hands5[i] = generateRandomHand(5)
	}
	
	fmt.Printf("Evaluating %d random 5-card hands...\n", numHands)
	startTime := time.Now()
	
	for i := 0; i < numHands; i++ {
		_ = evaluator.Evaluate(hands5[i])
	}
	
	elapsedTime5 := time.Since(startTime)
	handsPerSecond5 := float64(numHands) / elapsedTime5.Seconds()
	
	fmt.Printf("5-card: Evaluated %d hands in %v\n", numHands, elapsedTime5)
	fmt.Printf("5-card hands per second: %.0f\n\n", handsPerSecond5)
	
	// Test 7-card hands
	hands7 := make([]CardSet, numHands)
	for i := 0; i < numHands; i++ {
		hands7[i] = generateRandomHand(7)
	}
	
	fmt.Printf("Evaluating %d random 7-card hands...\n", numHands)
	startTime = time.Now()
	
	for i := 0; i < numHands; i++ {
		_ = evaluator.Evaluate(hands7[i])
	}
	
	elapsedTime7 := time.Since(startTime)
	handsPerSecond7 := float64(numHands) / elapsedTime7.Seconds()
	
	fmt.Printf("7-card: Evaluated %d hands in %v\n", numHands, elapsedTime7)
	fmt.Printf("7-card hands per second: %.0f\n\n", handsPerSecond7)
	
	// Compare
	fmt.Printf("Performance comparison:\n")
	fmt.Printf("5-card: %.0f hands/second\n", handsPerSecond5)
	fmt.Printf("7-card: %.0f hands/second\n", handsPerSecond7)
	fmt.Printf("Ratio: 5-card is %.2fx faster than 7-card\n", handsPerSecond5/handsPerSecond7)
}

func generateRandomHand(numCards int) CardSet {
	cards := []string{"2c", "2d", "2h", "2s", "3c", "3d", "3h", "3s", "4c", "4d", "4h", "4s", "5c", "5d", "5h", "5s", "6c", "6d", "6h", "6s", "7c", "7d", "7h", "7s", "8c", "8d", "8h", "8s", "9c", "9d", "9h", "9s", "Tc", "Td", "Th", "Ts", "Jc", "Jd", "Jh", "Js", "Qc", "Qd", "Qh", "Qs", "Kc", "Kd", "Kh", "Ks", "Ac", "Ad", "Ah", "As"}
	
	// Shuffle deck
	for i := len(cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	
	// Create hand string
	handString := ""
	for i := 0; i < numCards; i++ {
		handString += cards[i]
	}
	
	return NewHand(handString)
}