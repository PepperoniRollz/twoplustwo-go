package twoplustwogo

import (
	"math/rand"
	"testing"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(42))
}

func TestSimulator(t *testing.T) {
	players := make([]CardSet, 3)
	deck := NewDeck()
	deck.Shuffle(r)
	for i := 0; i < len(players); i++ {
		players[i] = FromCards(deck.Deal(2))
	}

	board := FromCards(deck.Deal(3))

	equityEval := EquityEvaluator(players, board)
	equityEval.Print()
}

func benchmarkSimulator(numPlayers, numCardsOnBoard int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		players := make([]CardSet, numPlayers)
		var board CardSet
		deck := NewDeck()
		deck.Shuffle(r)
		for i := 0; i < len(players); i++ {
			players[i] = FromCards(deck.Deal(2))
		}
		board.AddCards(FromCards(deck.Deal(numCardsOnBoard)))
		EquityEvaluator(players, board)
	}
}

// benchmark is run with 2 players, 0 cards on board, 2 players, 3 cards on board, and 2 players, 4 cards on board, etc.
// all the way from 2 players no board to 9 players, 4 cards on board
// simulator % go test -bench=. -benchtime=5s
// goos: darwin
// goarch: arm64
// pkg: github.com/pepperonirollz/twoplustwo-go/simulator
// BenchmarkSim2b0-8             15         342798742 ns/op
// BenchmarkSim2b3-8          33968            176250 ns/op
// BenchmarkSim2b4-8         511425             10761 ns/op
// BenchmarkSim3b0-8             14         362119405 ns/op
// BenchmarkSim3b3-8          27306            222225 ns/op
// BenchmarkSim3b4-8         396966             14088 ns/op
// BenchmarkSim4b0-8             15         363521303 ns/op
// BenchmarkSim4b3-8          23808            250625 ns/op
// BenchmarkSim4b4-8         352440             16304 ns/op
// BenchmarkSim5b0-8             16         324823914 ns/op
// BenchmarkSim5b3-8          21984            273330 ns/op
// BenchmarkSim5b4-8         305589             18888 ns/op
// BenchmarkSim6b0-8             20         279911279 ns/op
// BenchmarkSim6b3-8          22348            268196 ns/op
// BenchmarkSim6b4-8         284391             20267 ns/op
// BenchmarkSim7b0-8             22         241396574 ns/op
// BenchmarkSim7b3-8          20960            285361 ns/op
// BenchmarkSim7b4-8         276052             21315 ns/op
// BenchmarkSim8b0-8             28         197516629 ns/op
// BenchmarkSim8b3-8          20712            286667 ns/op
// BenchmarkSim8b4-8         254686             23262 ns/op
// BenchmarkSim9b0-8             33         164578612 ns/op
// BenchmarkSim9b3-8          21414            279930 ns/op
// BenchmarkSim9b4-8         240248             24841 ns/op
func BenchmarkSim2b0(b *testing.B) { benchmarkSimulator(2, 0, b) }
func BenchmarkSim2b3(b *testing.B) { benchmarkSimulator(2, 3, b) }
func BenchmarkSim2b4(b *testing.B) { benchmarkSimulator(2, 4, b) }

/////

func BenchmarkSim3b0(b *testing.B) { benchmarkSimulator(3, 0, b) }
func BenchmarkSim3b3(b *testing.B) { benchmarkSimulator(3, 3, b) }
func BenchmarkSim3b4(b *testing.B) { benchmarkSimulator(3, 4, b) }

/////

func BenchmarkSim4b0(b *testing.B) { benchmarkSimulator(4, 0, b) }
func BenchmarkSim4b3(b *testing.B) { benchmarkSimulator(4, 3, b) }
func BenchmarkSim4b4(b *testing.B) { benchmarkSimulator(4, 4, b) }

/////

func BenchmarkSim5b0(b *testing.B) { benchmarkSimulator(5, 0, b) }
func BenchmarkSim5b3(b *testing.B) { benchmarkSimulator(5, 3, b) }
func BenchmarkSim5b4(b *testing.B) { benchmarkSimulator(5, 4, b) }

/////

func BenchmarkSim6b0(b *testing.B) { benchmarkSimulator(6, 0, b) }
func BenchmarkSim6b3(b *testing.B) { benchmarkSimulator(6, 3, b) }
func BenchmarkSim6b4(b *testing.B) { benchmarkSimulator(6, 4, b) }

/////

func BenchmarkSim7b0(b *testing.B) { benchmarkSimulator(7, 0, b) }
func BenchmarkSim7b3(b *testing.B) { benchmarkSimulator(7, 3, b) }
func BenchmarkSim7b4(b *testing.B) { benchmarkSimulator(7, 4, b) }

/////

func BenchmarkSim8b0(b *testing.B) { benchmarkSimulator(8, 0, b) }
func BenchmarkSim8b3(b *testing.B) { benchmarkSimulator(8, 3, b) }
func BenchmarkSim8b4(b *testing.B) { benchmarkSimulator(8, 4, b) }

/////

func BenchmarkSim9b0(b *testing.B) { benchmarkSimulator(9, 0, b) }
func BenchmarkSim9b3(b *testing.B) { benchmarkSimulator(9, 3, b) }
func BenchmarkSim9b4(b *testing.B) { benchmarkSimulator(9, 4, b) }

// ----------------------------------------------------------------------------------
//  Hole Cards | Board      |Equity     | TieEquity  | Wins       | Losses     | Ties       |
// ----------------------------------------------------------------------------------
//  2♡J♢       | 8♡A♡Q♢     |0.0033     | 0.0554     | 3          | 850        | 50         |
//  8♢6♡       | 8♡A♡Q♢     |0.1495     | 0.4086     | 135        | 399        | 369        |
//  8♣J♣       | 8♡A♡Q♢     |0.3832     | 0.4640     | 346        | 138        | 419        |
// -----------------------------------------------------------------------------------
