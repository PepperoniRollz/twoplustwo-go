package generator

import (
	"fmt"
	"math"
)

func swap(i int, j int, workcards []int) {
	if workcards[i] < workcards[j] {
		workcards[i] ^= workcards[j]
		workcards[j] ^= workcards[i]
		workcards[i] ^= workcards[j]
	}
}

func MakeId(IDin int64, newcard int, numCards *int) int64 {
	var ID int64 = 0
	var suitCount [4 + 1]int
	var rankCount [13 + 1]int
	workcards := make([]int, 8)
	var cardnum int
	var getout int = 0

	for cardnum = 0; cardnum < 6; cardnum++ {
		workcards[cardnum+1] = (int)((IDin >> (8 * cardnum)) & 0xff) // leave the 0 hole for new card
	}

	// my cards are 2c = 1, 2d = 2  ... As = 52
	newcard-- // make 0 based!

	workcards[0] = (((newcard >> 2) + 1) << 4) + (newcard & 3) + 1 // add next card formats card to rrrr00ss

	for *numCards = 0; workcards[*numCards] != 0; *numCards++ {
		suitCount[workcards[*numCards]&0xf]++      // need to see if suit is significant
		rankCount[(workcards[*numCards]>>4)&0xf]++ // and rank to be sure we don't have 4!

		if *numCards != 0 {
			if workcards[0] == workcards[*numCards] { // can't have the same card twice
				getout = 1 // if so need to get out after counting numcards
			}
		}
	}
	if getout != 0 {
		return 0
	}
	var needsuited int = *numCards - 2

	if *numCards > 4 {
		for rank := 1; rank < 14; rank++ {
			if rankCount[rank] > 4 {
				return 0
			}
		}
	}

	if needsuited > 1 {
		for cardnum := 0; cardnum < *numCards; cardnum++ {
			if suitCount[workcards[cardnum]&0xf] < needsuited {
				workcards[cardnum] &= 0xf0 // mask out suit
			}
		}
	}

	swap(0, 4, workcards)
	swap(1, 5, workcards)
	swap(2, 6, workcards)
	swap(0, 2, workcards)
	swap(1, 3, workcards)
	swap(4, 6, workcards)
	swap(2, 4, workcards)
	swap(3, 5, workcards)
	swap(0, 1, workcards)
	swap(2, 3, workcards)
	swap(4, 5, workcards)
	swap(1, 4, workcards)
	swap(3, 6, workcards)
	swap(1, 2, workcards)
	swap(3, 4, workcards)
	swap(5, 6, workcards)

	ID = int64(workcards[0]) +
		(int64(workcards[1] << 8)) +
		(int64(workcards[2] << 16)) +
		(int64(workcards[3] << 24)) +
		(int64(workcards[4] << 32)) +
		(int64(workcards[5] << 40)) +
		(int64(workcards[6] << 48))

	return ID
}
func SaveId(ID int64, IDs []int64, numIds *int, maxId *int64) int {

	if ID == 0 {
		return 0
	}
	if ID >= *maxId {
		if ID > *maxId {
			IDs[*numIds] = ID
			*numIds++
			*maxId = ID
		}
		return *numIds - 1
	}

	var low int = 0
	var high int = *numIds - 1
	var testval int64
	var holdtest int

	for high-low > 1 {
		holdtest = (high + low + 1) / 2
		testval = IDs[holdtest] - ID

		if testval > 0 {
			high = holdtest
		} else if testval < 0 {
			low = holdtest
		} else {
			return holdtest
		}
	}
	copy(IDs[high+1:], IDs[high:*numIds])
	IDs[high] = ID
	*numIds++
	return high
}

func DoEval(IDin int64) int {
	var handrank int = 0
	var cardnum int
	var workcard int
	var rank int
	var suit int
	var mainsuit int = 20
	var suititerator int = 0
	var holdrank int
	var workcards [8]int
	var holdcards [8]int
	var numevalcards int = 0

	if IDin != 0 {
		for cardnum = 0; cardnum < 7; cardnum++ {
			holdcards[cardnum] = (int)((IDin >> (8 * cardnum)) & 0xff)
			if holdcards[cardnum] == 0 {
				break
			}
			numevalcards++
			if suit == holdcards[cardnum]&0xf {
				mainsuit = suit
			}

		}

		for cardnum = 0; cardnum < numevalcards; cardnum++ {
			workcard = holdcards[cardnum]

			rank = (workcard >> 4) - 1
			suit = workcard & 0xf

			if suit == 0 {
				suit = suititerator
				suititerator++
				if suititerator == 5 {
					suititerator = 1
				}
				if suit == mainsuit {
					suit = suititerator
					suititerator++
					if suititerator == 5 {
						suititerator = 1
					}
				}
			}
			workcards[cardnum] = PRIMES[rank] | (rank << 8) | (1 << (suit + 11)) | 1<<(16+rank)
		}
		switch numevalcards {
		case 5:
			holdrank = eval5HandFast(uint32(workcards[0]), uint32(workcards[1]), uint32(workcards[2]), uint32(workcards[3]), uint32(workcards[4]))
			break

		case 6:
			holdrank = eval5HandFast(uint32(workcards[0]), uint32(workcards[1]), uint32(workcards[2]), uint32(workcards[3]), uint32(workcards[4]))
			holdrank = int(math.Min(float64(holdrank), float64(eval5HandFast(uint32(workcards[0]), uint32(workcards[1]), uint32(workcards[2]), uint32(workcards[3]), uint32(workcards[5])))))
			holdrank = int(math.Min(float64(holdrank), float64(eval5HandFast(uint32(workcards[0]), uint32(workcards[1]), uint32(workcards[2]), uint32(workcards[4]), uint32(workcards[5])))))
			holdrank = int(math.Min(float64(holdrank), float64(eval5HandFast(uint32(workcards[0]), uint32(workcards[1]), uint32(workcards[3]), uint32(workcards[4]), uint32(workcards[5])))))
			holdrank = int(math.Min(float64(holdrank), float64(eval5HandFast(uint32(workcards[0]), uint32(workcards[2]), uint32(workcards[3]), uint32(workcards[4]), uint32(workcards[5])))))
			holdrank = int(math.Min(float64(holdrank), float64(eval5HandFast(uint32(workcards[1]), uint32(workcards[2]), uint32(workcards[3]), uint32(workcards[4]), uint32(workcards[5])))))
			break
		case 7:
			holdrank = eval7hand(workcards)
			break
		default:
			fmt.Printf("Problem with numcards = %d!!\n", cardnum)

		}
		handrank = 7463 - holdrank //now the worst hand = 1
		if handrank < 1278 {
			handrank = handrank - 0 + 4096*1
		} else if handrank < 4138 {
			handrank = handrank - 1277 + 4096*2
		} else if handrank < 4996 {
			handrank = handrank - 4137 + 4096*3
		} else if handrank < 5854 {
			handrank = handrank - 4995 + 4096*4
		} else if handrank < 5864 {
			handrank = handrank - 5853 + 4096*5
		} else if handrank < 7141 {
			handrank = handrank - 5863 + 4096*6
		} else if handrank < 7297 {
			handrank = handrank - 7140 + 4096*7
		} else if handrank < 7453 {
			handrank = handrank - 7296 + 4096*8
		} else {
			handrank = handrank - 7452 + 4096*9
		}
	}
	return handrank
}
