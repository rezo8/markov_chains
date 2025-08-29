package stateMachines

import (
	"markov_chains/types"
)

func NewCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Tails: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateCounter: map[types.CoinFlipState]int{},
		Description:  "Fair Coin Flip Chain. Each flip has 50 percent chance to flip.\n",
	}
}

func NewEvenlyBiasedCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.95, types.Tails: 0.05, types.Start: 0.0},
			types.Tails: {types.Heads: 0.05, types.Tails: 0.95, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateCounter: map[types.CoinFlipState]int{},
		Description:  "Evenly Biased Coin Flip Chain. Each flip has a 95 percent chance to stay on original flip.\n",
	}
}

func NewUnEvenlyBiasedCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		TransitionMatrix: map[types.CoinFlipState]map[types.CoinFlipState]float64{
			types.Start: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
			types.Heads: {types.Heads: 0.95, types.Tails: 0.05, types.Start: 0.0},
			types.Tails: {types.Heads: 0.5, types.Tails: 0.5, types.Start: 0.0},
		},
		CurrentState: types.Start,
		StateCounter: map[types.CoinFlipState]int{},
		Description:  "Unevenly Biased Coin Flip Chain. Tails flip has a 50 percent chance to flip while Heads has a 5%.\n",
	}
}
