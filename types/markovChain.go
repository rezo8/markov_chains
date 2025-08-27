// markov_chain.go
// Defines the MarkovChain struct

package types

type MarkovChain struct {
	States           []State
	TransitionMatrix map[State]map[State]float64
	CurrentState     State
}
