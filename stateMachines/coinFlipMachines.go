package stateMachines

import (
	"markov_chains/helpers"
	"markov_chains/types"
)

func NewCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		MatrixChain: NewMatrixChain(
			helpers.Matrix{
				types.CoinFlip_Heads: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5},
				types.CoinFlip_Tails: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5},
			},
			"Fair Coin Flip Chain. Each flip has 50 percent chance to flip.\n",
			nil,
		),
	}
}

func NewEvenlyBiasedCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		MatrixChain: NewMatrixChain(
			helpers.Matrix{
				types.CoinFlip_Heads: {types.CoinFlip_Heads: 0.95, types.CoinFlip_Tails: 0.05},
				types.CoinFlip_Tails: {types.CoinFlip_Heads: 0.05, types.CoinFlip_Tails: 0.95},
			},
			"Evenly Biased Coin Flip Chain. Each flip has a 95 percent chance to stay on original flip.\n",
			nil,
		),
	}
}

func NewUnEvenlyBiasedCoinFlipChain() *CoinFlipChain {
	return &CoinFlipChain{
		MatrixChain: NewMatrixChain(
			helpers.Matrix{
				types.CoinFlip_Heads: {types.CoinFlip_Heads: 0.95, types.CoinFlip_Tails: 0.05},
				types.CoinFlip_Tails: {types.CoinFlip_Heads: 0.5, types.CoinFlip_Tails: 0.5},
			},
			"Unevenly Biased Coin Flip Chain. Tails flip has a 50 percent chance to flip while Heads has a 5%.\n",
			nil,
		),
	}
}
