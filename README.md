# twoplustwo-go

A fast, modern Go implementation of the TwoPlusTwo poker hand evaluator with smart caching and flexible configuration.

Creates a 248MB lookup table capable of evaluating 5-7 card poker hands as fast as your machine can do array lookups (~1 microsecond per hand).

## Installation

```bash
go get github.com/pepperonirollz/twoplustwo-go
```

## Quick Start

```bash
# Try the example first
go run cmd/example/main.go
```

```go
package main

import (
    "fmt"
    twoplustwo "github.com/pepperonirollz/twoplustwo-go"
)

func main() {
    // Create a poker hand - Royal Flush in Spades
    cards := twoplustwo.NewHand("AsKsQsJsTs") // Ace, King, Queen, Jack, Ten of spades
    
    // Evaluate the hand (uses smart caching - auto-generates if needed)
    result := twoplustwo.Evaluate(cards)
    
    fmt.Printf("Hand type: %v\n", result.HandCategory) // 9 = Straight Flush
    fmt.Printf("Hand strength: %v\n", result.Value)    // Higher = stronger
}
```

## Features

- **🚀 Ultra-fast evaluation**: ~1 microsecond per hand after initialization
- **🔄 Smart caching**: Auto-generates lookup table once, reuses across projects
- **⚙️ Flexible configuration**: Custom file paths, progress callbacks, parallel generation
- **🔒 Thread-safe**: Concurrent evaluations supported
- **📦 Zero dependencies**: Pure Go implementation
- **🔙 Backwards compatible**: Drop-in replacement for existing code

## Advanced Usage

### Custom Configuration

```go
evaluator, err := twoplustwo.NewEvaluator(twoplustwo.Config{
    CacheDir: "/custom/cache",     // Custom cache directory
    Verbose:  true,                // Show generation progress
    OnProgress: func(stage string, percent float64) {
        log.Printf("Poker engine: %s %.1f%%", stage, percent)
    },
    ParallelGeneration: true,      // Use multiple CPU cores
})
```

### Production Deployment

```go
// For servers/containers - specify exact file location
evaluator, err := twoplustwo.NewEvaluatorWithPath("/app/data/HandRanks.dat", twoplustwo.Config{
    Verbose: false, // Silent operation
})
```

### Card Format

Cards use simple 2-character strings:
- **Ranks**: `2, 3, 4, 5, 6, 7, 8, 9, T, J, Q, K, A`  
- **Suits**: `c` (clubs), `d` (diamonds), `h` (hearts), `s` (spades)`

Examples:
- `"As"` = Ace of Spades
- `"Kh"` = King of Hearts
- `"AsKsQsJsTs"` = Royal Flush (5 cards)
- `"AsKsQsJsTs7h2d"` = 7-card hand (finds best 5)

## Performance

### Benchmarks

Tested on **MacBook Pro M4 Pro (24GB RAM, macOS 15.5)** - July 2025:

#### Hand Evaluation Speed
- **5-card hands**: 24.1 million hands/second (41.6ns per hand)
- **7-card hands**: 17.2 million hands/second (58.2ns per hand)  
- **Individual lookup**: 13.5ns per hand (zero allocations)

#### Full Enumeration Performance
- **All 133,784,560 possible 7-card hands**: 201ms (665M hands/second)
- **Statistical validation**: Perfect match with known poker probabilities
- **Memory**: Zero allocations during evaluation

#### Equity Calculation Benchmarks
- **2 players, no board**: 161ms per simulation
- **2 players, 4 board cards**: 6.1μs per simulation  
- **9 players, 4 board cards**: 13.6μs per simulation
- **Concurrent evaluation**: Thread-safe, scales with CPU cores

#### System Requirements
- **Setup time**: 3-5 minutes (generates lookup table once)
- **Load time**: ~200ms (loads 248MB cached file)
- **Memory usage**: ~248MB (lookup table in RAM)
- **Disk usage**: ~248MB (cached HandRanks.dat file)

## Hand Categories

| Value | Hand Type |
|-------|-----------|
| 0 | Invalid |
| 1 | High Card |
| 2 | One Pair |
| 3 | Two Pair |
| 4 | Three of a Kind |
| 5 | Straight |
| 6 | Flush |
| 7 | Full House |
| 8 | Four of a Kind |
| 9 | Straight Flush |

## API Reference

### Core Functions

- `Evaluate(cards CardSet) HandEvaluation` - Evaluate a poker hand
- `CompareHands(hand1, hand2 HandEvaluation) int` - Compare two hands (-1, 0, 1)
- `Best5(cards CardSet) CardSet` - Find best 5 cards from 6-7 card hand

### Initialization

- `MustNewEvaluator() *Evaluator` - Simple setup (panics on error)
- `NewEvaluator(config Config) (*Evaluator, error)` - Configurable setup
- `NewEvaluatorWithPath(path string, config Config) (*Evaluator, error)` - Custom file path

## License

Original TwoPlusTwo algorithm. Go implementation improvements and modernization.