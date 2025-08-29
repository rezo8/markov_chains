package stateMachines

import (
	"markov_chains/helpers"
	"markov_chains/types"
)

func NewCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		TransitionMatrix: helpers.Matrix{
			types.CoinFlip_Start: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5, types.CoinFlip_Start: 0.0},
			types.CoinFlip_Heads: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5, types.CoinFlip_Start: 0.0},
			types.CoinFlip_Tails: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5, types.CoinFlip_Start: 0.0},
		},
		CurrentState: types.CoinFlip_Start,
		StateCounter: map[types.State]int{},
		Description:  "Fair Coin Flip Chain. Each flip has 50 percent chance to flip.\n",
	}
}

func NewEvenlyBiasedCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		TransitionMatrix: map[types.State]map[types.State]float64{
			types.CoinFlip_Start: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5, types.CoinFlip_Start: 0.0},
			types.CoinFlip_Heads: {types.CoinFlip_Heads: 0.95, types.CoinFlip_Tails: 0.05, types.CoinFlip_Start: 0.0},
			types.CoinFlip_Tails: {types.CoinFlip_Heads: 0.05, types.CoinFlip_Tails: 0.95, types.CoinFlip_Start: 0.0},
		},
		CurrentState: types.CoinFlip_Start,
		StateCounter: map[types.State]int{},
		Description:  "Evenly Biased Coin Flip Chain. Each flip has a 95 percent chance to stay on original flip.\n",
	}
}

func NewUnEvenlyBiasedCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		TransitionMatrix: map[types.State]map[types.State]float64{
			types.CoinFlip_Start: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5, types.CoinFlip_Start: 0.0},
			types.CoinFlip_Heads: {types.CoinFlip_Heads: 0.95, types.CoinFlip_Tails: 0.05, types.CoinFlip_Start: 0.0},
			types.CoinFlip_Tails: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5, types.CoinFlip_Start: 0.0},
		},
		CurrentState: types.CoinFlip_Start,
		StateCounter: map[types.State]int{},
		Description:  "Unevenly Biased Coin Flip Chain. Tails flip has a 50 percent chance to flip while Heads has a 5%.\n",
	}
}
