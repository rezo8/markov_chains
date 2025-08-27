package types

import "math/rand"

type CoinFlipState State

const (
	Start CoinFlipState = "Start"
	Heads CoinFlipState = "Heads"
	Tails CoinFlipState = "Tails"
)

type CoinFlipChain struct {
	States           []CoinFlipState
	TransitionMatrix map[CoinFlipState]map[CoinFlipState]float64
	CurrentState     CoinFlipState
	StateHistory     map[CoinFlipState]int
}

func (c *CoinFlipChain) Step() {
	randVal := rand.Float64()
	nextState := c.CurrentState
	for state, prob := range c.TransitionMatrix[c.CurrentState] {
		if randVal < prob {
			nextState = state
			break
		}
		randVal -= prob
	}
	c.CurrentState = nextState
	c.StateHistory[nextState]++
}
