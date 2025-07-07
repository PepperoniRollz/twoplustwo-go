package twoplustwogo

import (
	"fmt"
	"testing"
)

func TestEquityExamples(t *testing.T) {
	fmt.Println("🃏 Poker Equity Examples")
	fmt.Println("========================")

	// PREFLOP EQUITIES
	fmt.Println("\n📋 PREFLOP EQUITIES:")
	
	// Preflop 1: AsKs vs 2c2d (classic close matchup)
	fmt.Println("\n1. AsKs vs 2c2d (preflop):")
	testEquity("AsKs", "2c2d", "")
	
	// Preflop 2: AAh vs KQs (premium vs strong)
	fmt.Println("\n2. AhAc vs KsQs (preflop):")
	testEquity("AhAc", "KsQs", "")

	// FLOP EQUITIES
	fmt.Println("\n📋 FLOP EQUITIES:")
	
	// Flop 1: AsKs vs 7c2c on 5s6d7h (overcards + flush draw vs top pair)
	fmt.Println("\n3. AsKs vs 7c2c on flop 5s6d7h:")
	testEquity("AsKs", "7c2c", "5s6d7h")
	
	// Flop 2: AhAc vs 8h9h on As2d3h (set vs flush draw)
	fmt.Println("\n4. AhAc vs 8h9h on flop As2d3h:")
	testEquity("AhAc", "8h9h", "As2d3h")

	// TURN EQUITIES
	fmt.Println("\n📋 TURN EQUITIES:")
	
	// Turn 1: AsKs vs 7c2c on 5s6d7h4c (missed draws vs made pair)
	fmt.Println("\n5. AsKs vs 7c2c on turn 5s6d7h4c:")
	testEquity("AsKs", "7c2c", "5s6d7h4c")
	
	// Turn 2: QhJh vs AdKc on Qd7h2s9h (pair + flush draw vs overcards)
	fmt.Println("\n6. QhJh vs AdKc on turn Qd7h2s9h:")
	testEquity("QhJh", "AdKc", "Qd7h2s9h")
}

func testEquity(hand1Str, hand2Str, boardStr string) {
	hand1 := NewHand(hand1Str)
	hand2 := NewHand(hand2Str)
	
	var board CardSet
	if boardStr != "" {
		board = NewHand(boardStr)
	}
	
	holeCards := []CardSet{hand1, hand2}
	equityEval := EvaluateEquity(holeCards, board)
	
	fmt.Printf("   %s: %.2f%%\n", hand1Str, equityEval.Equities[0]*100)
	fmt.Printf("   %s: %.2f%%\n", hand2Str, equityEval.Equities[1]*100)
	
	winner := hand1Str
	margin := equityEval.Equities[0] - equityEval.Equities[1]
	if equityEval.Equities[1] > equityEval.Equities[0] {
		winner = hand2Str
		margin = equityEval.Equities[1] - equityEval.Equities[0]
	}
	fmt.Printf("   Winner: %s (+%.2f%%)\n", winner, margin*100)
}