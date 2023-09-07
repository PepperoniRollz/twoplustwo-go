package evaluator

import (
	"encoding/binary"
	"fmt"
	"os"

	card "github.com/pepperonirollz/twoplustwo-go/card"
)

type Evaluator struct {
	HR []int64
}

func (e *Evaluator) GetHandValue(pCards card.CardSet) HandEvaluation {
	var p int64 = 53
	size := len(pCards.Cards)
	if size < 5 {
		panic("Not enough cards to evaluate hand.")
	}
	if size > 7 {
		panic("Too many cards to evaluate hand.")
	}
	for i := 0; i < size; i++ {
		p = e.HR[p+int64(pCards.Cards[i].Value)]
	}

	if size == 5 || size == 6 {
		p = e.HR[p]
	}

	return newHandEval(p, pCards.Cards)
}

func (e *Evaluator) CompareHands(hand1 HandEvaluation, hand2 HandEvaluation) int {

	if hand1.P > hand2.P {
		return 1
	} else if hand1.P < hand2.P {
		return -1
	} else {
		return 0
	}
}

func NewEvaluator(pathToHandRanks string) Evaluator {
	file, err := os.Open(pathToHandRanks)
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
