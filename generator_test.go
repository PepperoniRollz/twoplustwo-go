package twoplustwogo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	evaluator := NewEvaluator("../HandRanks.dat")

	fmt.Println("Initialization complete.")

	handTypeSum := EnumerateAll7CardHands(evaluator.HR)

	assert.Equal(t, 133784560, handTypeSum[1]+handTypeSum[2]+handTypeSum[3]+handTypeSum[4]+handTypeSum[5]+handTypeSum[6]+handTypeSum[7]+handTypeSum[8]+handTypeSum[9], "correct total number of 7 card hands")
	assert.Equal(t, 0, handTypeSum[0], "correct number of invalid hands")
	assert.Equal(t, 23294460, handTypeSum[1], "correct number of high card hands")
	assert.Equal(t, 58627800, handTypeSum[2], "correct number of one pair hands")
	assert.Equal(t, 31433400, handTypeSum[3], "correct number of two pair hands")
	assert.Equal(t, 6461620, handTypeSum[4], "correct number of trips hands")
	assert.Equal(t, 6180020, handTypeSum[5], "correct number of straight hands")
	assert.Equal(t, 4047644, handTypeSum[6], "correct number of flush hands")
	assert.Equal(t, 3473184, handTypeSum[7], "correct number of full house hands")
	assert.Equal(t, 224848, handTypeSum[8], "correct number of quads hands")
	assert.Equal(t, 41584, handTypeSum[9], "correct number of straight flush hands")

}

func EnumerateAll7CardHands(HR []int64) []int {
	var u0, u1, u2, u3, u4, u5 int64
	var c0, c1, c2, c3, c4, c5, c6 int64
	handTypeSum := make([]int, 10)
	count := 0

	fmt.Println("Enumerating and evaluating all 133,784,560 possible 7-card poker hands...")

	startTime := time.Now()

	for c0 = 1; c0 < 47; c0++ {
		u0 = HR[53+c0]
		for c1 = c0 + 1; c1 < 48; c1++ {
			u1 = HR[u0+c1]
			for c2 = c1 + 1; c2 < 49; c2++ {
				u2 = HR[u1+c2]
				for c3 = c2 + 1; c3 < 50; c3++ {
					u3 = HR[u2+c3]
					for c4 = c3 + 1; c4 < 51; c4++ {
						u4 = HR[u3+c4]
						for c5 = c4 + 1; c5 < 52; c5++ {
							u5 = HR[u4+c5]
							for c6 = c5 + 1; c6 < 53; c6++ {
								handTypeSum[HR[u5+c6]>>12]++
								count++
							}
						}
					}
				}
			}
		}
	}

	elapsedTime := time.Since(startTime)

	fmt.Printf("BAD:              %d\n", handTypeSum[0])
	fmt.Printf("High Card:        %d\n", handTypeSum[1])
	fmt.Printf("One Pair:         %d\n", handTypeSum[2])
	fmt.Printf("Two Pair:         %d\n", handTypeSum[3])
	fmt.Printf("Trips:            %d\n", handTypeSum[4])
	fmt.Printf("Straight:         %d\n", handTypeSum[5])
	fmt.Printf("Flush:            %d\n", handTypeSum[6])
	fmt.Printf("Full House:       %d\n", handTypeSum[7])
	fmt.Printf("Quads:            %d\n", handTypeSum[8])
	fmt.Printf("Straight Flush:   %d\n", handTypeSum[9])

	testCount := 0
	for _, val := range handTypeSum {
		testCount += val
	}
	if testCount != count || count != 133784560 || handTypeSum[0] != 0 {
		fmt.Println("\nERROR!\nERROR!\nERROR!")
		return []int{0}
	}

	fmt.Printf("\nEnumerated %d hands in %v!\n", count, elapsedTime)
	return handTypeSum
}
