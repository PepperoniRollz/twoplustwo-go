package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	gen "github.com/pepperonirollz/twoplustwo-go/generator"
)

func main() {
	var IdSlot, card, count int = 0, 0, 0
	var ID int64

	handSumType := make([]int, 10)

	var IdNum int

	IDs := make([]int64, 612978)
	var handRanks = [...]string{
		"BAD!!",
		"High Card",
		"Pair",
		"Two Pair",
		"Three of a Kind",
		"Straight",
		"Flush",
		"Full House",
		"Four of a Kind",
		"Straight Flush",
	}

	HR := make([]int, 32487834)

	var numIds *int = new(int)
	*numIds = 1

	var maxHR *int = new(int)
	var maxId *int64 = new(int64)
	var numCards *int = new(int)

	fmt.Println("...Starting...Getting Cards!")

	for IdNum = 0; IDs[IdNum] != 0 || IdNum == 0; IdNum++ {

		for card = 1; card < 53; card++ {
			ID = gen.MakeId(IDs[IdNum], card, numCards)
			if *numCards < 7 {
				gen.SaveId(ID, IDs, numIds, maxId)
			}
			fmt.Printf("\rID - %d", IdNum) // Just to show progress, counting up to 612976.

		}
	}

	fmt.Printf("\nSetting Handranks!\n")

	for IdNum = 0; IDs[IdNum] != 0 || IdNum == 0; IdNum++ {

		for card = 1; card < 53; card++ {
			ID = gen.MakeId(IDs[IdNum], card, numCards)
			if *numCards < 7 {
				IdSlot = gen.SaveId(ID, IDs, numIds, maxId)*53 + 53

			} else {
				IdSlot = gen.DoEval(ID)
			}

			*maxHR = IdNum*53 + card + 53
			HR[*maxHR] = IdSlot
			fmt.Printf("\rID - %d", IdNum) // Just to show progress, counting up to 612976.
		}
		if *numCards == 6 || *numCards == 7 {
			HR[IdNum*53+53] = gen.DoEval(IDs[IdNum]) // this puts the above handrank into the array

		}

	}

	fmt.Printf("\nNumber IDs = %d\nmaxHR = %d\n", *numIds, *maxHR)

	var c0, c1, c2, c3, c4, c5, c6 int
	var u0, u1, u2, u3, u4, u5 int

	timings := time.Now() // Start a timer

	for c0 = 1; c0 < 53; c0++ {
		u0 = HR[53+c0]
		for c1 = c0 + 1; c1 < 53; c1++ {
			u1 = HR[u0+c1]
			for c2 = c1 + 1; c2 < 53; c2++ {
				u2 = HR[u1+c2]
				for c3 = c2 + 1; c3 < 53; c3++ {
					u3 = HR[u2+c3]
					for c4 = c3 + 1; c4 < 53; c4++ {
						u4 = HR[u3+c4]
						for c5 = c4 + 1; c5 < 53; c5++ {
							u5 = HR[u4+c5]
							for c6 = c5 + 1; c6 < 53; c6++ {
								handSumType[HR[u5+c6]>>12]++
								count++
							}
						}
					}
				}
			}
		}
	}

	elapsed := time.Since(timings)
	fmt.Printf("Elapsed time: %v\n", elapsed)

	for i := 0; i <= 9; i++ {
		fmt.Printf("\n%16s = %d", handRanks[i], handSumType[i])
	}

	fmt.Printf("\nTotal Hands = %d\n", count)

	fout, err := os.Create("HandRanks.dat")
	if err != nil {
		fmt.Println("Problem creating the Output File!")
		return
	}
	defer fout.Close()

	byteArray := make([]byte, len(HR)*8)

	for i, v := range HR {
		binary.LittleEndian.PutUint64(byteArray[i*8:], uint64(v))
	}

	_, err = fout.Write(byteArray[:])
	if err != nil {
		fmt.Println("Problem writing to the Output File!")
		return
	}

}
