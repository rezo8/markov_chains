package stateMachines

import (
	"fmt"
	"markov_chains/helpers"
	"markov_chains/types"
	"math"
	"math/rand"
	"time"
)

type HigherOrderMatrixChain struct {
	Matrix      helpers.HigherOrderMatrix
	History     helpers.HigherOrderState
	Description string
	K           int

	rng         *rand.Rand
	stateLog    []types.State
	stateCounts map[types.State]int
	totalSteps  int
}

func NewHigherOrderMatrixChain(matrix helpers.HigherOrderMatrix, initial helpers.HigherOrderState, description string) *HigherOrderMatrixChain {
	if len(initial.States) == 0 {
		panic("initial history cannot be empty")
	}
	return &HigherOrderMatrixChain{
		Matrix:      matrix,
		Description: description,
		History:     initial,
		K:           len(initial.States),
		rng:         rand.New(rand.NewSource(time.Now().UnixNano())),
		stateLog:    []types.State{},
		stateCounts: map[types.State]int{},
	}
}

func (c *HigherOrderMatrixChain) step() types.State {
	row, ok := c.Matrix[c.History.StateKey()]
	if !ok {
		panic(fmt.Sprintf("unknown history: %v", c.History))
	}

	r := c.rng.Float64()
	cum := 0.0
	var next types.State
	for st, p := range row {
		cum += p
		if r <= cum {
			next = st
			break
		}
	}

	// fallback for rounding
	if next == "" {
		for st := range row {
			next = st
			break
		}
	}

	// Update history
	c.History.States = append(c.History.States[1:], next)
	c.stateLog = append(c.stateLog, next)
	c.stateCounts[next]++
	c.totalSteps++

	if len(c.History.States) > c.K {
		c.History.States = c.History.States[len(c.History.States)-c.K:]
	}

	return next
}

// Run simulation for n steps
func (m *HigherOrderMatrixChain) RunSimulation(steps int) {
	m.reset()
	for i := 0; i < steps; i++ {
		m.step()
	}
	m.printResults()
}

// Reset internal state
func (m *HigherOrderMatrixChain) reset() {
	m.stateCounts = map[types.State]int{}
	m.stateLog = []types.State{}
	m.totalSteps = 0
}

// Print results
func (m *HigherOrderMatrixChain) printResults() {
	fmt.Printf("=== %s Results ===\n", m.Description)
	fmt.Printf("Total Steps: %d\n", len(m.stateLog))

	keys := make([]types.State, 0, len(m.stateCounts))
	for s := range m.stateCounts {
		keys = append(keys, s)
	}

	for _, state := range keys {
		count := m.stateCounts[state]
		fmt.Printf("%s: %d (%.2f%%)\n", state, count, float64(count)*100/float64(len(m.stateLog)))
	}
	// fmt.Printf("Longest Streak: %d\n", m.longestStreak)
	// fmt.Printf("First Passage Time: %d\n", m.firstPassage)
	fmt.Printf("Entropy: %.4f\n", m.Entropy())
	fmt.Println()
}

// TODO Fix this.
func (c *HigherOrderMatrixChain) PredictNthState(n int) helpers.HigherOrderMatrix {
	// Need a way to map from key to HigherOrderstate

	keys := make([]helpers.HigherOrderState, 0, len(c.Matrix))
	for k := range c.Matrix {
		// Assuming k is already a HigherOrderState, otherwise convert appropriately
		keys = append(keys, helpers.ToHigherOrderState(k))
	}
	return helpers.MatPowHO(c.Matrix, n, keys)
}

func (c *HigherOrderMatrixChain) Entropy() float64 {
	row, ok := c.Matrix[c.History.StateKey()]
	if !ok {
		return 0
	}
	h := 0.0
	for _, p := range row {
		if p > 0 {
			h -= p * math.Log2(p)
		}
	}
	return h
}
