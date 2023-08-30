package generator

import (
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"testing"
	"time"
)

func InitTheEvaluator() []int64 {
	file, err := os.Open("../HandRanks.dat")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	HR := make([]int64, 32487834)
	if err := binary.Read(file, binary.LittleEndian, &HR); err != nil {
		fmt.Println("Error reading HR data:", err)

	}
	return HR
}

func GetHandValue(pCards []int, HR []int64) int64 {

	var p int64 = 53
	size := len(pCards)
	for i := 0; i < size; i++ {
		p = HR[p+int64(pCards[i])]
	}

	if size == 5 || size == 6 {
		p = HR[p]
	}

	return p
}

func TestGenerator(t *testing.T) {
	var Hr []int64 = InitTheEvaluator()
	fmt.Println("Initialization complete.")

	pCards := []int{1, 5, 9, 13, 25, 52} // Example card values
	result := GetHandValue(pCards, Hr)

	handCategory := result >> 12
	rankWithinCategory := result & 0x00000FFF

	sort.Ints(pCards)
	var hand []string
	for _, v := range pCards {
		hand = append(hand, CARD_MAP[v])
	}
	fmt.Println(hand)
	fmt.Println("Hand value:", result)

	fmt.Println("Hand category:", HAND_TYPES[int(handCategory)])

	fmt.Println("rank within category:", rankWithinCategory)

	EnumerateAll7CardHands(Hr)
}

func EnumerateAll7CardHands(HR []int64) {
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

	// Perform sanity checks.. make sure numbers are where they should be
	testCount := 0
	for _, val := range handTypeSum {
		testCount += val
	}
	if testCount != count || count != 133784560 || handTypeSum[0] != 0 {
		fmt.Println("\nERROR!\nERROR!\nERROR!")
		return
	}

	fmt.Printf("\nEnumerated %d hands in %v!\n", count, elapsedTime)
}
