package twoplustwogo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestNewAPIFunctionality tests the new flexible API
func TestNewAPIFunctionality(t *testing.T) {
	// Test the new API with custom config
	evaluator, err := NewEvaluator(Config{
		Verbose: false, // Silent for testing
	})
	if err != nil {
		t.Fatalf("Failed to create evaluator: %v", err)
	}

	// Test basic evaluation with new API
	cards := &CardSet{}
	cards.NewCardSet("AsKsQsJsTs") // Royal flush
	
	result := evaluator.Evaluate(*cards)
	if result.HandCategory != 9 {
		t.Errorf("Expected straight flush (9), got %v", result.HandCategory)
	}
}

// TestConcurrentHandEvaluation tests thread safety for poker solvers
func TestConcurrentHandEvaluation(t *testing.T) {
	const numGoroutines = 10
	const handsPerGoroutine = 1000
	
	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			
			// Each goroutine evaluates many hands
			for j := 0; j < handsPerGoroutine; j++ {
				cards := &CardSet{}
				cards.NewCardSet("AsKsQsJsTs") // Same hand for consistency
				
				result := Evaluate(*cards)
				if result.HandCategory != 9 {
					errors <- fmt.Errorf("Goroutine %d, hand %d: expected 9, got %v", i, j, result.HandCategory)
					return
				}
			}
		}(i)
	}
	
	wg.Wait()
	close(errors)
	
	// Check for any errors
	for err := range errors {
		t.Error(err)
	}
	
	t.Logf("Successfully evaluated %d hands across %d goroutines", numGoroutines*handsPerGoroutine, numGoroutines)
}

// TestHandComparisonPerformance tests speed of hand comparisons for solvers
func TestHandComparisonPerformance(t *testing.T) {
	// Pre-create test hands
	hands := make([]CardSet, 1000)
	for i := 0; i < len(hands); i++ {
		cards := &CardSet{}
		// Create various hands for testing
		if i%4 == 0 {
			cards.NewCardSet("AsKsQsJsTs") // Royal flush
		} else if i%4 == 1 {
			cards.NewCardSet("AhAcAdAs7d") // Four aces
		} else if i%4 == 2 {
			cards.NewCardSet("AhKhQhJhTh") // Royal flush hearts
		} else {
			cards.NewCardSet("2c3c4c5c6c") // Straight flush low
		}
		hands[i] = *cards
	}
	
	// Benchmark hand evaluations
	start := time.Now()
	results := make([]HandEvaluation, len(hands))
	for i, hand := range hands {
		results[i] = Evaluate(hand)
	}
	evalTime := time.Since(start)
	
	// Benchmark hand comparisons
	start = time.Now()
	comparisons := 0
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			CompareHands(results[i], results[j])
			comparisons++
		}
	}
	compareTime := time.Since(start)
	
	avgEvalTime := evalTime / time.Duration(len(hands))
	avgCompareTime := compareTime / time.Duration(comparisons)
	
	t.Logf("Evaluated %d hands in %v (avg: %v per hand)", len(hands), evalTime, avgEvalTime)
	t.Logf("Performed %d comparisons in %v (avg: %v per comparison)", comparisons, compareTime, avgCompareTime)
	
	// Performance assertions for poker solver use
	if avgEvalTime > 10*time.Microsecond {
		t.Errorf("Hand evaluation too slow: %v (expected < 10μs)", avgEvalTime)
	}
	if avgCompareTime > 100*time.Nanosecond {
		t.Errorf("Hand comparison too slow: %v (expected < 100ns)", avgCompareTime)
	}
}

// TestEquityCalculation tests the equity evaluation for solver use
func TestEquityCalculation(t *testing.T) {
	// Test heads-up equity calculation
	holeCards := make([]CardSet, 2)
	
	// Player 1: Pocket Aces
	cards1 := &CardSet{}
	cards1.NewCardSet("AsAh")
	holeCards[0] = *cards1
	
	// Player 2: King Queen suited
	cards2 := &CardSet{}
	cards2.NewCardSet("KsQs")
	holeCards[1] = *cards2
	
	// Empty board (preflop)
	board := CardSet{}
	
	start := time.Now()
	equity := EvaluateEquity(holeCards, board)
	duration := time.Since(start)
	
	t.Logf("Equity calculation took: %v", duration)
	t.Logf("AA vs KQs equity: Player 1: %.3f, Player 2: %.3f", 
		equity.Equities[0], equity.Equities[1])
	
	// Pocket Aces should have significant equity advantage
	if equity.Equities[0] < 0.7 {
		t.Errorf("AA should have >70%% equity vs KQs, got %.3f", equity.Equities[0])
	}
	
	// Equities should sum to ~1.0 (allowing for ties)
	totalEquity := equity.Equities[0] + equity.Equities[1]
	if totalEquity < 0.95 || totalEquity > 1.05 {
		t.Errorf("Equities should sum to ~1.0, got %.3f", totalEquity)
	}
}

// TestBatchHandEvaluation tests evaluating many hands efficiently
func TestBatchHandEvaluation(t *testing.T) {
	// Create a range of hands to evaluate
	const numHands = 10000
	hands := make([]CardSet, numHands)
	
	// Generate random-ish hands for testing
	suits := []string{"c", "d", "h", "s"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	
	for i := 0; i < numHands; i++ {
		cards := &CardSet{}
		// Create 5-card hands with some variety
		handStr := ""
		for j := 0; j < 5; j++ {
			rank := ranks[(i*7+j)%len(ranks)]
			suit := suits[j%len(suits)]
			handStr += rank + suit
		}
		cards.NewCardSet(handStr)
		hands[i] = *cards
	}
	
	start := time.Now()
	
	// Evaluate all hands
	results := make([]HandEvaluation, numHands)
	for i, hand := range hands {
		results[i] = Evaluate(hand)
	}
	
	duration := time.Since(start)
	avgTime := duration / time.Duration(numHands)
	
	t.Logf("Evaluated %d hands in %v (avg: %v per hand)", numHands, duration, avgTime)
	
	// Should be very fast for poker solver use
	if avgTime > 5*time.Microsecond {
		t.Errorf("Batch evaluation too slow: %v per hand (expected < 5μs)", avgTime)
	}
	
	// Verify results are reasonable
	handTypes := make(map[int]int)
	for _, result := range results {
		handTypes[result.HandCategory]++
	}
	
	t.Logf("Hand type distribution: %v", handTypes)
}