package stateMachines

import (
	"fmt"
	"markov_chains/helpers"
	"markov_chains/types"
	"math"
	"math/rand"
)

type MatrixChain struct {
	TransitionMatrix helpers.Matrix
	Description      string

	possibleStates []types.State
	currentState   types.State
	stateCounter   map[types.State]int
	stateLog       []types.State
	longestStreak  int
	currentStreak  int
	firstPassage   int
	totalSteps     int
}

func NewMatrixChain(transitionMatrix helpers.Matrix, description string, startState *types.State) *MatrixChain {
	// Extract keys from the transition matrix to infer possible states
	possibleStates := make([]types.State, 0, len(transitionMatrix))
	for state := range transitionMatrix {
		possibleStates = append(possibleStates, state)
	}

	var currentState types.State // Default to uniform start state selection
	if startState == nil {
		randIndex := rand.Intn(len(possibleStates))
		currentState = possibleStates[randIndex]
	} else {
		currentState = *startState
	}

	return &MatrixChain{
		TransitionMatrix: transitionMatrix,
		possibleStates:   possibleStates,
		Description:      description,
		currentState:     currentState,
		longestStreak:    0,
		currentStreak:    0,
	}
}

func (m *MatrixChain) runSimulation(steps int) {
	m.reset()
	// Generate off of possible states
	stateCounts := make(map[types.State]int)
	for _, state := range m.possibleStates {
		stateCounts[state] = 0
	}

	for i := 0; i < steps; i++ {
		m.step()
		stateCounts[m.currentState]++
	}

	m.printResults()
}

func (m *MatrixChain) printResults() {
	fmt.Printf("=== %s Results ===\n", m.Description)
	fmt.Printf("Total Steps: %d\n", len(m.stateLog))

	// Generate report based on keys of stateCounter
	for state, count := range m.stateCounter {
		fmt.Printf("%s: %d (%.2f%%)\n", state, count, float64(count)*100/float64(len(m.stateLog)))
	}
	fmt.Printf("Longest Streak: %d\n", m.longestStreak)
	fmt.Printf("First Passage Time to Change State: %d\n", m.firstPassage)
	fmt.Printf("Entropy: %.4f\n", m.calculateEntropy())

	fmt.Println()
}

func (m *MatrixChain) reset() {
	m.stateCounter = map[types.State]int{}
	m.stateLog = []types.State{}
	m.longestStreak = 0
	m.totalSteps = 0
	m.firstPassage = 0
}

func (m *MatrixChain) step() {
	// Increment total steps
	m.totalSteps++

	randVal := rand.Float64()
	nextState := m.currentState

	for state, prob := range m.TransitionMatrix[m.currentState] {
		if randVal < prob {
			nextState = state
			break
		}
		randVal -= prob
	}

	if nextState == m.currentState {
		m.currentStreak++
	} else {
		if m.firstPassage == 0 {
			m.firstPassage = m.totalSteps
		}
		m.currentStreak = 1
	}

	if m.currentStreak > m.longestStreak {
		m.longestStreak = m.currentStreak
	}

	m.currentState = nextState
	m.stateCounter[nextState]++
	m.stateLog = append(m.stateLog, nextState)
}

func (m *MatrixChain) PredictNthState(n int) helpers.Matrix {
	return helpers.MatPow(m.TransitionMatrix, n, m.possibleStates)
}

func (m *MatrixChain) calculateEntropy() float64 {
	pArray := make([]float64, len(m.possibleStates))
	for i, state := range m.possibleStates {
		pArray[i] = float64(m.stateCounter[state]) / float64(m.totalSteps)
	}
	entropy := 0.0
	for _, p := range pArray {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}
