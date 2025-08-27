package main

import (
	"markov_chains/stateMachines"
	"markov_chains/types"
)

func NewCoinFlipChain() *stateMachines.CoinFlipChain {
	return &stateMachines.CoinFlipChain{
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Tails: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateCounter: map[types.CoinFlipState]int{},
	}
}

func NewEvenlyBiasedCoinFlipChain() *stateMachines.CoinFlipChain {
	return &stateMachines.CoinFlipChain{
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.95, types.Tails: 0.05, types.Start: 0.0},
			types.Tails: {types.Heads: 0.05, types.Tails: 0.95, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateCounter: map[types.CoinFlipState]int{},
	}
}

func NewUnEvenlyBiasedCoinFlipChain() *stateMachines.CoinFlipChain {
	return &stateMachines.CoinFlipChain{
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.95, types.Tails: 0.05, types.Start: 0.0},
			types.Tails: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateCounter: map[types.CoinFlipState]int{},
	}
}

func main() {
	const steps = 10000

	coinFlipChain := NewCoinFlipChain()
	evenlyBiasedCoinFlipChain := NewEvenlyBiasedCoinFlipChain()
	unEvenlyBiasedCoinFlipChain := NewUnEvenlyBiasedCoinFlipChain()

	countsFair := coinFlipChain.RunSimulation(steps)
	countsWeighted := evenlyBiasedCoinFlipChain.RunSimulation(steps)
	countsUnWeighted := unEvenlyBiasedCoinFlipChain.RunSimulation(steps)

	coinFlipChain.PrintResults("Fair Coin Flip Chain. Each flip has 50 percent chance to flip.\n", countsFair, steps)
	evenlyBiasedCoinFlipChain.PrintResults("Evenly Biased Coin Flip Chain. Each flip has a 95 percent chance to stay on original flip.\n", countsWeighted, steps)
	unEvenlyBiasedCoinFlipChain.PrintResults("Unevenly biased Coin Flip Chain. Tails flip has a 50 percent chance to flip while Heads has a 5%.\n", countsUnWeighted, steps)

	const visuallyFriendlySteps = 500
	coinFlipChain.RunSimulation(visuallyFriendlySteps)
	evenlyBiasedCoinFlipChain.RunSimulation(visuallyFriendlySteps)
	unEvenlyBiasedCoinFlipChain.RunSimulation(visuallyFriendlySteps)

	coinFlipChain.PlotStateSequence("fair_coin_flip_chain.png")
	evenlyBiasedCoinFlipChain.PlotStateSequence("evenly_biased_coin_flip_chain.png")
	unEvenlyBiasedCoinFlipChain.PlotStateSequence("unevenly_biased_coin_flip_chain.png")

	coinFlipChain.Reset()
	evenlyBiasedCoinFlipChain.Reset()
	unEvenlyBiasedCoinFlipChain.Reset()
}
