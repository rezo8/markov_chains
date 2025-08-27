package main

import (
	"fmt"
	"markov_chains/types"
)

func NewCoinFlipChain() *types.CoinFlipChain {
	return &types.CoinFlipChain{
		States: []types.CoinFlipState{types.Heads, types.Tails},
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Heads: {types.Heads: 0.5, types.Tails: 0.5},
			types.Tails: {types.Heads: 0.5, types.Tails: 0.5},
		},
		CurrentState: types.Heads,
		StateHistory: map[types.CoinFlipState]int{},
	}
}

func NewEvenlyBiasedCoinFlipChain() *types.CoinFlipChain {
	initialStates := []types.CoinFlipState{types.Heads, types.Tails}
	return &types.CoinFlipChain{
		States: initialStates,
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.95, types.Tails: 0.05, types.Start: 0.0},
			types.Tails: {types.Heads: 0.05, types.Tails: 0.95, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateHistory: map[types.CoinFlipState]int{},
	}
}

func NewUnEvenlyBiasedCoinFlipChain() *types.CoinFlipChain {
	initialStates := []types.CoinFlipState{types.Heads, types.Tails}
	return &types.CoinFlipChain{
		States: initialStates,
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.95, types.Tails: 0.05, types.Start: 0.0},
			types.Tails: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateHistory: map[types.CoinFlipState]int{},
	}
}

func runSimulation(chain *types.CoinFlipChain, steps int) map[types.CoinFlipState]int {
	stateCounts := map[types.CoinFlipState]int{
		types.Heads: 0,
		types.Tails: 0,
	}
	for i := 0; i < steps; i++ {
		chain.Step()
		stateCounts[chain.CurrentState]++
	}
	return stateCounts
}

func printResults(name string, counts map[types.CoinFlipState]int, steps int) {
	fmt.Printf("=== %s Results ===\n", name)
	fmt.Printf("Total Steps: %d\n", steps)
	fmt.Printf("Heads: %d (%.2f%%)\n", counts[types.Heads], float64(counts[types.Heads])*100/float64(steps))
	fmt.Printf("Tails: %d (%.2f%%)\n", counts[types.Tails], float64(counts[types.Tails])*100/float64(steps))
	fmt.Println()
}

func main() {
	const steps = 10000

	coinFlipChain := NewCoinFlipChain()
	evenlyBiasedCoinFlipChain := NewEvenlyBiasedCoinFlipChain()
	unEvenlyBiasedCoinFlipChain := NewUnEvenlyBiasedCoinFlipChain()

	countsFair := runSimulation(coinFlipChain, steps)
	countsWeighted := runSimulation(evenlyBiasedCoinFlipChain, steps)
	countsUnWeighted := runSimulation(unEvenlyBiasedCoinFlipChain, steps)

	printResults("Fair Coin Flip Chain. Each flip has 50 percent chance to flip.\n", countsFair, steps)
	printResults("Evenly Biased Coin Flip Chain. Each flip has a 95 percent chance to stay on original flip.\n", countsWeighted, steps)
	printResults("Unevenly biased Coin Flip Chain. Tails flip has a 50 percent chance to flip while Heads has a 5%.\n", countsUnWeighted, steps)
}
