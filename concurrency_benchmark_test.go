package twoplustwogo

import (
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestConcurrencyPerformance(t *testing.T) {
	evaluator := getDefaultEvaluator()
	numHands := 1000000
	
	// Generate test hands
	hands := make([]CardSet, numHands)
	for i := 0; i < numHands; i++ {
		hands[i] = generateRandomHand(7)
	}
	
	// Test 1: Single goroutine
	t.Log("Testing single goroutine...")
	start := time.Now()
	for i := 0; i < numHands; i++ {
		_ = evaluator.Evaluate(hands[i])
	}
	singleTime := time.Since(start)
	singleRate := float64(numHands) / singleTime.Seconds()
	
	// Test 2: Multiple goroutines (number of CPU cores)
	numCores := runtime.NumCPU()
	t.Logf("Testing %d goroutines (CPU cores)...", numCores)
	
	start = time.Now()
	var wg sync.WaitGroup
	handsPerGoroutine := numHands / numCores
	
	for i := 0; i < numCores; i++ {
		wg.Add(1)
		go func(startIdx int) {
			defer wg.Done()
			endIdx := startIdx + handsPerGoroutine
			if endIdx > numHands {
				endIdx = numHands
			}
			for j := startIdx; j < endIdx; j++ {
				_ = evaluator.Evaluate(hands[j])
			}
		}(i * handsPerGoroutine)
	}
	wg.Wait()
	multiTime := time.Since(start)
	multiRate := float64(numHands) / multiTime.Seconds()
	
	// Test 3: Excessive goroutines (2x CPU cores)
	numExcess := numCores * 2
	t.Logf("Testing %d goroutines (2x CPU cores)...", numExcess)
	
	start = time.Now()
	handsPerGoroutine = numHands / numExcess
	
	for i := 0; i < numExcess; i++ {
		wg.Add(1)
		go func(startIdx int) {
			defer wg.Done()
			endIdx := startIdx + handsPerGoroutine
			if endIdx > numHands {
				endIdx = numHands
			}
			for j := startIdx; j < endIdx; j++ {
				_ = evaluator.Evaluate(hands[j])
			}
		}(i * handsPerGoroutine)
	}
	wg.Wait()
	excessTime := time.Since(start)
	excessRate := float64(numHands) / excessTime.Seconds()
	
	// Results
	t.Logf("\nPerformance Results:")
	t.Logf("Single goroutine:     %v (%.0f hands/sec)", singleTime, singleRate)
	t.Logf("%d goroutines:        %v (%.0f hands/sec) - %.2fx speedup", numCores, multiTime, multiRate, multiRate/singleRate)
	t.Logf("%d goroutines:        %v (%.0f hands/sec) - %.2fx speedup", numExcess, excessTime, excessRate, excessRate/singleRate)
	
	t.Logf("\nCPU cores available: %d", numCores)
	
	if multiRate <= singleRate {
		t.Log("❌ Multiple goroutines did NOT improve performance")
	} else {
		t.Log("✅ Multiple goroutines improved performance")
	}
}