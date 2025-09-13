package stateMachines

import (
	"markov_chains/helpers"
	"markov_chains/types"
)

func NewCoinFlipChain() *CoinFlipHigherOrderChain {

	matrix := helpers.HigherOrderMatrix{
		helpers.HigherOrderState{States: []types.State{types.CoinFlip_Heads}}.StateKey(): {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5},
		helpers.HigherOrderState{States: []types.State{types.CoinFlip_Tails}}.StateKey(): {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5},
	}
	return &CoinFlipHigherOrderChain{
		HigherOrderMatrixChain: NewHigherOrderMatrixChain(
			matrix,
			helpers.HigherOrderState{States: []types.State{types.CoinFlip_Heads}},
			"Fair Coin Flip Chain. Each flip has 50 percent chance to flip.\n",
			1,
		),
	}
}

func NewEvenlyBiasedCoinFlipChain() *CoinFlipHigherOrderChain {

	matrix := helpers.HigherOrderMatrix{
		helpers.HigherOrderState{States: []types.State{types.CoinFlip_Heads}}.StateKey(): {types.CoinFlip_Heads: 0.95, types.CoinFlip_Tails: 0.05},
		helpers.HigherOrderState{States: []types.State{types.CoinFlip_Tails}}.StateKey(): {types.CoinFlip_Heads: 0.05, types.CoinFlip_Tails: 0.95},
	}

	return &CoinFlipHigherOrderChain{
		HigherOrderMatrixChain: NewHigherOrderMatrixChain(
			matrix,
			helpers.HigherOrderState{States: []types.State{types.CoinFlip_Heads}},
			"Evenly Biased Coin Flip Chain. Each flip has a 95 percent chance to stay on original flip.\n",
			1,
		),
	}
}

func NewUnEvenlyBiasedCoinFlipChain() *CoinFlipHigherOrderChain {
	return &CoinFlipHigherOrderChain{
		HigherOrderMatrixChain: NewHigherOrderMatrixChain(
			helpers.HigherOrderMatrix{
				helpers.HigherOrderState{States: []types.State{types.CoinFlip_Heads}}.StateKey(): {types.CoinFlip_Heads: 0.95, types.CoinFlip_Tails: 0.05},
				helpers.HigherOrderState{States: []types.State{types.CoinFlip_Tails}}.StateKey(): {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5},
			},
			helpers.HigherOrderState{States: []types.State{types.CoinFlip_Heads}},
			"Unevenly Biased Coin Flip Chain. Tails flip has a 50 percent chance to flip while Heads has a 5%.\n",
			1,
		),
	}
}
