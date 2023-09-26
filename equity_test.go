package twoplustwogo

import (
	"fmt"
	"math/rand"
	"testing"
)

var r *rand.Rand

func TestSimulator(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	players := make([]CardSet, 3)
	deck := NewDeck()
	deck.Shuffle(r)
	for i := 0; i < len(players); i++ {
		players[i].AddCards(deck.Deal(2))
	}
	board := deck.Deal(3)
	fmt.Println("Board:", board)
	equityEval := EvaluateEquity(players, board)
	equityEval.Print()
}

func benchmarkEquity(numPlayers, numCardsOnBoard int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		players := make([]CardSet, numPlayers)
		var board CardSet
		deck := NewDeck()
		deck.Shuffle(r)
		for i := 0; i < len(players); i++ {
			players[i].AddCards(deck.Deal(2))
		}
		board.AddCards(deck.Deal(numCardsOnBoard))
		EvaluateEquity(players, board)
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

func BenchmarkEquity2Players0Board(b *testing.B) { benchmarkEquity(2, 0, b) }
func BenchMarkEquity2Players3Board(b *testing.B) { benchmarkEquity(2, 3, b) }
func BenchmarkEquity2Players4Board(b *testing.B) { benchmarkEquity(2, 4, b) }

/////

func BenchmarkEquity3Players0Board(b *testing.B) { benchmarkEquity(3, 0, b) }
func BenchmarkEquity3Players3Board(b *testing.B) { benchmarkEquity(3, 3, b) }
func BenchmarkEquity3Players4Board(b *testing.B) { benchmarkEquity(3, 4, b) }

/////

func BenchmarkEquity4Players0Board(b *testing.B) { benchmarkEquity(4, 0, b) }
func BenchmarkEquity4Players3Board(b *testing.B) { benchmarkEquity(4, 3, b) }
func BenchmarkEquity4Players4Board(b *testing.B) { benchmarkEquity(4, 4, b) }

/////

func BenchmarkEquity5Players0Board(b *testing.B) { benchmarkEquity(5, 0, b) }
func BenchMarkEquity5Players3Board(b *testing.B) { benchmarkEquity(5, 3, b) }
func BenchmarkEquity5Players4Board(b *testing.B) { benchmarkEquity(5, 4, b) }

/////

func BenchmarkEquity6Players0Board(b *testing.B) { benchmarkEquity(6, 0, b) }
func BenchMarkEquity6Players3Board(b *testing.B) { benchmarkEquity(6, 3, b) }
func BenchmarkEquity6Players4Board(b *testing.B) { benchmarkEquity(6, 4, b) }

/////

func BenchmarkEquity7Players0Board(b *testing.B) { benchmarkEquity(7, 0, b) }
func BenchMarkEquity7Players3Board(b *testing.B) { benchmarkEquity(7, 3, b) }
func BenchmarkEquity7Players4Board(b *testing.B) { benchmarkEquity(7, 4, b) }

/////

func BenchmarkEquity8Players0Board(b *testing.B) { benchmarkEquity(8, 0, b) }
func BenchMarkEquity8Players3Board(b *testing.B) { benchmarkEquity(8, 3, b) }
func BenchmarkEquity8Players4Board(b *testing.B) { benchmarkEquity(8, 4, b) }

/////

func BenchmarkEquity9Players0Board(b *testing.B) { benchmarkEquity(9, 0, b) }
func BenchMarkEquity9Players3Board(b *testing.B) { benchmarkEquity(9, 3, b) }
func BenchmarkEquity9Players4Board(b *testing.B) { benchmarkEquity(9, 4, b) }

// ----------------------------------------------------------------------------------
//  Hole Cards | Board      |Equity     | TieEquity  | Wins       | Losses     | Ties       |
// ----------------------------------------------------------------------------------
//  2♡J♢       | 8♡A♡Q♢     |0.0033     | 0.0554     | 3          | 850        | 50         |
//  8♢6♡       | 8♡A♡Q♢     |0.1495     | 0.4086     | 135        | 399        | 369        |
//  8♣J♣       | 8♡A♡Q♢     |0.3832     | 0.4640     | 346        | 138        | 419        |
// -----------------------------------------------------------------------------------
