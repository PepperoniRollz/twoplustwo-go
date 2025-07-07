package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	twoplustwo "github.com/pepperonirollz/twoplustwo-go"
)

func main() {
	fmt.Println("🃏 TwoPlusTwo Go Hand Evaluator Example")
	fmt.Println("=====================================")

	// Example 1: Basic hand evaluation
	fmt.Println("\n1. Basic Hand Evaluation:")

	// Royal Flush in Spades
	royalFlush := twoplustwo.NewHand("AsKsQsJsTs")
	result := twoplustwo.Evaluate(royalFlush)
	fmt.Printf("   Royal Flush: %s -> Category: %d, Value: %d\n",
		"As Ks Qs Js Ts", result.HandCategory, result.Value)

	// Pair of Aces
	pairAces := twoplustwo.NewHand("AsAh7c5d2s")
	result = twoplustwo.Evaluate(pairAces)
	fmt.Printf("   Pair of Aces: %s -> Category: %d, Value: %d\n",
		"As Ah 7c 5d 2s", result.HandCategory, result.Value)

	// High Card
	highCard := twoplustwo.NewHand("Kh9c7d5s2h")
	result = twoplustwo.Evaluate(highCard)
	fmt.Printf("   High Card: %s -> Category: %d, Value: %d\n",
		"Kh 9c 7d 5s 2h", result.HandCategory, result.Value)

	// Example 2: Performance demonstration
	fmt.Println("\n2. Performance Test:")

	// Generate random hands for testing
	hands := generateTestHands(100000)
	fmt.Printf("   Generated %d random hands\n", len(hands))

	// Single-threaded evaluation
	start := time.Now()
	for _, hand := range hands {
		_ = twoplustwo.Evaluate(hand)
	}
	singleTime := time.Since(start)
	singleRate := float64(len(hands)) / singleTime.Seconds()

	fmt.Printf("   Single-threaded: %v (%.0f hands/sec)\n", singleTime, singleRate)

	// Multi-threaded evaluation
	numCores := runtime.NumCPU()
	start = time.Now()

	var wg sync.WaitGroup
	handsPerCore := len(hands) / numCores

	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go func(startIdx int) {
			defer wg.Done()
			endIdx := startIdx + handsPerCore
			if endIdx > len(hands) {
				endIdx = len(hands)
			}
			for j := startIdx; j < endIdx; j++ {
				_ = twoplustwo.Evaluate(hands[j])
			}
		}(i * handsPerCore)
	}
	wg.Wait()

	multiTime := time.Since(start)
	multiRate := float64(len(hands)) / multiTime.Seconds()
	speedup := multiRate / singleRate

	fmt.Printf("   Multi-threaded (%d cores): %v (%.0f hands/sec) - %.1fx speedup\n",
		numCores, multiTime, multiRate, speedup)

	// Example 3: Hand comparison
	fmt.Println("\n3. Hand Comparison:")

	hand1 := twoplustwo.NewHand("AsAhKcQdJs") // Pair of Aces
	hand2 := twoplustwo.NewHand("KsKhAcQdJs") // Pair of Kings

	eval1 := twoplustwo.Evaluate(hand1)
	eval2 := twoplustwo.Evaluate(hand2)

	comparison := twoplustwo.CompareHands(eval1, eval2)

	fmt.Printf("   Hand 1: %s (Category: %d, Value: %d)\n",
		"As Ah Kc Qd Js", eval1.HandCategory, eval1.Value)
	fmt.Printf("   Hand 2: %s (Category: %d, Value: %d)\n",
		"Ks Kh Ac Qd Js", eval2.HandCategory, eval2.Value)

	switch comparison {
	case 1:
		fmt.Printf("   Result: Hand 1 wins!\n")
	case -1:
		fmt.Printf("   Result: Hand 2 wins!\n")
	case 0:
		fmt.Printf("   Result: Tie!\n")
	}

	// Example 4: Equity calculations
	fmt.Println("\n4. Equity Calculations:")

	// Preflop equity
	fmt.Println("\n   Preflop: AsKs vs 2c2d")
	playerHands := []twoplustwo.CardSet{
		twoplustwo.NewHand("AsKs"),
		twoplustwo.NewHand("2c2d"),
	}
	board := twoplustwo.CardSet{} // Empty board (preflop)

	equity := twoplustwo.EvaluateEquity(playerHands, board)
	fmt.Printf("   AsKs: %.1f%%\n", equity.Equities[0]*100)
	fmt.Printf("   2c2d: %.1f%%\n", equity.Equities[1]*100)

	// Flop equity
	fmt.Println("\n   Flop: AsKs vs 7c2c on 5s6d7h")
	playerHands = []twoplustwo.CardSet{
		twoplustwo.NewHand("AsKs"),
		twoplustwo.NewHand("7c2c"),
	}
	board = twoplustwo.NewHand("5s6d7h") // Flop board

	equity = twoplustwo.EvaluateEquity(playerHands, board)
	fmt.Printf("   AsKs: %.1f%% (overcards + flush draw)\n", equity.Equities[0]*100)
	fmt.Printf("   7c2c: %.1f%% (top pair)\n", equity.Equities[1]*100)
}

func generateTestHands(count int) []twoplustwo.CardSet {
	hands := make([]twoplustwo.CardSet, count)

	// Generate some variety of hands for testing
	handPatterns := []string{
		"AsKsQsJsTs", // Royal flush
		"AhAcAdAs7d", // Four aces
		"KhKcKd7s7h", // Full house
		"AhKcQdJsTc", // High card
		"AhAcKdQsJs", // Pair of aces
		"KhKcQdQsJs", // Two pair
	}

	for i := 0; i < count; i++ {
		pattern := handPatterns[i%len(handPatterns)]
		hands[i] = twoplustwo.NewHand(pattern)
	}

	return hands
}

