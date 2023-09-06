package evaluator

import (
	"encoding/binary"
	"fmt"
	"os"

	d "github.com/pepperonirollz/twoplustwo-go/deck"
)

type Evaluator struct {
	HR []int64
}

func (e *Evaluator) GetHandValue(pCards []d.Card) HandEvaluation {
	var p int64 = 53
	size := len(pCards)
	if size < 5 {
		panic("Not enough cards to evaluate hand.")
	}
	if size > 7 {
		panic("Too many cards to evaluate hand.")
	}
	for i := 0; i < size; i++ {
		p = e.HR[p+int64(pCards[i].Value)]
	}

	if size == 5 || size == 6 {
		p = e.HR[p]
	}

	return NewHand(p, pCards)
}

func (e *Evaluator) CompareHands(hand1 HandEvaluation, hand2 HandEvaluation) int {
	if hand1.HandCategory > hand2.HandCategory {
		return 1
	} else if hand1.HandCategory < hand2.HandCategory {
		return -1
	} else {
		if hand1.RankWithinCategory > hand2.RankWithinCategory {
			return 1
		} else if hand1.RankWithinCategory < hand2.RankWithinCategory {
			return -1
		} else {
			return 0
		}
	}
}

func NewEvaluator() Evaluator {
	file, err := os.Open("../HandRanks.dat")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	HR := make([]int64, 32487834)
	if err := binary.Read(file, binary.LittleEndian, &HR); err != nil {
		fmt.Println("Error reading HR data:", err)

	}

	return Evaluator{HR: HR}

}
