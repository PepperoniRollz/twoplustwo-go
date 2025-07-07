package twoplustwogo

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	defaultEvaluator *Evaluator
	defaultOnce      sync.Once
)

// Config holds configuration options for the evaluator
type Config struct {
	CacheDir           string
	FilePath           string
	OnProgress         func(stage string, percent float64)
	Verbose            bool
	ParallelGeneration bool
}

// getDefaultEvaluator returns the default evaluator instance, creating it if necessary
func getDefaultEvaluator() *Evaluator {
	defaultOnce.Do(func() {
		var err error
		defaultEvaluator, err = NewEvaluator(Config{
			Verbose: true, // Show progress for better UX during auto-generation
		})
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize default evaluator: %v", err))
		}
	})
	return defaultEvaluator
}

type Evaluator struct {
	HR []int64
}

// Evaluate evaluates a poker hand using the default evaluator
func Evaluate(pCards CardSet) HandEvaluation {
	return getDefaultEvaluator().Evaluate(pCards)
}

// Evaluate evaluates a poker hand using this evaluator instance
func (e *Evaluator) Evaluate(pCards CardSet) HandEvaluation {
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

	return newHandEval(p, pCards)
}

func CompareHands(hand1 HandEvaluation, hand2 HandEvaluation) int {

	if hand1.Value > hand2.Value {
		return 1
	} else if hand1.Value < hand2.Value {
		return -1
	} else {
		return 0
	}
}

// NewEvaluator creates a new evaluator with the given configuration
func NewEvaluator(config Config) (*Evaluator, error) {
	filePath := config.FilePath
	if filePath == "" {
		// First check for HandRanks.dat in current directory for backwards compatibility
		if _, err := os.Stat("HandRanks.dat"); err == nil {
			filePath = "HandRanks.dat"
		} else {
			// Fall back to cache system
			var err error
			filePath, err = getDefaultCachePath(config.CacheDir)
			if err != nil {
				return nil, fmt.Errorf("failed to get cache path: %w", err)
			}
		}
	}

	return NewEvaluatorWithPath(filePath, config)
}

// NewEvaluatorWithPath creates a new evaluator using the specified file path
func NewEvaluatorWithPath(pathToHandRanks string, config Config) (*Evaluator, error) {
	if config.Verbose {
		fmt.Printf("Loading HandRanks from: %s\n", pathToHandRanks)
	}

	file, err := os.Open(pathToHandRanks)
	if err != nil {
		if config.Verbose {
			fmt.Println("HandRanks.dat not found, generating...")
		}
		genConfig := GeneratorConfig{
			Verbose:            config.Verbose,
			OnProgress:         config.OnProgress,
			ParallelGeneration: config.ParallelGeneration,
		}
		if err := GenerateToFile(pathToHandRanks, genConfig); err != nil {
			return nil, fmt.Errorf("failed to generate HandRanks.dat: %w", err)
		}
		return NewEvaluatorWithPath(pathToHandRanks, config)
	}
	defer file.Close()

	// The file was written as int64, so read as int64 and convert
	hrData := make([]int64, 32487834)
	if err := binary.Read(file, binary.LittleEndian, &hrData); err != nil {
		return nil, fmt.Errorf("error reading HandRanks data: %w", err)
	}
	
	// Convert to int64 slice for the evaluator
	HR := hrData

	return &Evaluator{HR: HR}, nil
}

// MustNewEvaluator creates a new evaluator with default configuration, panicking on error
func MustNewEvaluator() *Evaluator {
	e, err := NewEvaluator(Config{})
	if err != nil {
		panic(err)
	}
	return e
}

// MustNewEvaluatorWithPath creates a new evaluator with the specified path, panicking on error
func MustNewEvaluatorWithPath(pathToHandRanks string) *Evaluator {
	e, err := NewEvaluatorWithPath(pathToHandRanks, Config{})
	if err != nil {
		panic(err)
	}
	return e
}

// Legacy function for backwards compatibility
func NewEvaluatorLegacy(pathToHandRanks string) Evaluator {
	e, err := NewEvaluatorWithPath(pathToHandRanks, Config{Verbose: true})
	if err != nil {
		panic(err)
	}
	return *e
}

// getDefaultCachePath returns the default cache path for HandRanks.dat
func getDefaultCachePath(customCacheDir string) (string, error) {
	var cacheDir string
	if customCacheDir != "" {
		cacheDir = customCacheDir
	} else {
		userCacheDir, err := os.UserCacheDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user cache directory: %w", err)
		}
		cacheDir = filepath.Join(userCacheDir, "twoplustwo-go")
	}

	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %w", err)
	}

	// Create a version hash for the algorithm
	versionHash := getAlgorithmHash()
	fileName := fmt.Sprintf("HandRanks-v%s.dat", versionHash[:8])
	return filepath.Join(cacheDir, fileName), nil
}

// getAlgorithmHash returns a hash representing the current algorithm version
func getAlgorithmHash() string {
	// This should be updated whenever the algorithm changes
	algorithmVersion := "twoplustwo-go-v1.0.0"
	hash := sha256.Sum256([]byte(algorithmVersion))
	return fmt.Sprintf("%x", hash)
}

// Best5 finds the best 5-card hand from the given cards using the default evaluator
func Best5(cards CardSet) CardSet {
	return getDefaultEvaluator().Best5(cards)
}

// Best5 finds the best 5-card hand from the given cards using this evaluator instance
func (e *Evaluator) Best5(cards CardSet) CardSet {
	if cards.Length() == 6 {
		var bestScore int64
		var bestI = 0
		for i := 0; i < 6; i++ {
			temp := cards
			temp.RemoveCard(cards.Get(i))
			score := e.Evaluate(temp).Value
			if score > bestScore {
				bestScore = score
				bestI = i
			}
		}
		cards.RemoveCard(cards.Get(bestI))
		return cards
	}

	if cards.Length() == 7 {
		var bestScore int64
		var bestI = 0
		for i := 0; i < 7; i++ {
			temp := cards
			temp.RemoveCard(cards.Get(i))
			score := e.Evaluate5(temp).Value
			if score > bestScore {
				bestScore = score
				bestI = i
			}
		}
		cards.RemoveCard(cards.Get(bestI))
		return cards
	}

	var best CardSet
	var bestHandEval int64 = -1
	combos := GenerateCombos(cards, 5)
	for i := 0; i < len(combos); i++ {
		handValue := e.Evaluate(combos[i])
		if handValue.Value > bestHandEval {
			best = combos[i]
		}
	}
	return best
}

// Evaluate5 evaluates a hand, finding the best 5 cards if more than 5 are provided
func Evaluate5(cards CardSet) HandEvaluation {
	return getDefaultEvaluator().Evaluate5(cards)
}

// Evaluate5 evaluates a hand, finding the best 5 cards if more than 5 are provided
func (e *Evaluator) Evaluate5(cards CardSet) HandEvaluation {
	if cards.Length() == 5 {
		return e.Evaluate(cards)
	}
	fiveBest := e.Best5(cards)
	return e.Evaluate(fiveBest)
}
