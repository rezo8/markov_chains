package helpers

import (
	"markov_chains/types"
	"strings"
)

// HigherOrderState = tuple of length K
type HigherOrderState struct {
	States []types.State
}

func (hos HigherOrderState) StateKey() string {
	rawStrings := make([]string, len(hos.States))
	for i, state := range hos.States {
		rawStrings[i] = string(state)
	}
	return strings.Join(rawStrings, "|")
}

func ToHigherOrderState(s string) HigherOrderState {
	parts := strings.Split(s, "|")
	states := make([]types.State, len(parts))
	for i, p := range parts {
		states[i] = types.State(p)
	}
	return HigherOrderState{States: states}
}

// HigherOrderMatrix: each history → distribution over NEXT single states.
type HigherOrderMatrix map[string]map[types.State]float64

// HigherOrderMatrixIdentity builds a uniform baseline distribution.
func HigherOrderMatrixIdentity(order []HigherOrderState) HigherOrderMatrix {
	stateSet := make(map[types.State]bool)
	for _, s := range order {
		for _, state := range s.States {
			stateSet[state] = true
		}
	}

	I := make(HigherOrderMatrix, len(order))
	for _, s := range order {
		dist := make(map[types.State]float64)
		totalStates := float64(len(stateSet))
		for st := range stateSet {
			dist[st] = 1.0 / totalStates // Uniform distribution
		}
		I[s.StateKey()] = dist
	}
	return I
}

// MatMulHO multiplies two HigherOrderMatrices
func MatMulHO(A, B HigherOrderMatrix, order []HigherOrderState) HigherOrderMatrix {
	C := make(HigherOrderMatrix, len(order))

	// Iterate over all rows (i) in the order
	for _, i := range order {
		iKey := i.StateKey()
		C[iKey] = make(map[types.State]float64)

		// Iterate over all columns (j) in the order
		for _, j := range order {
			jKey := j.StateKey()
			sum := 0.0
			for _, k := range order {
				kKey := k.StateKey()
				if probA, existsA := A[iKey][types.State(kKey)]; existsA {
					if probB, existsB := B[kKey][types.State(jKey)]; existsB {
						sum += probA * probB
					}
				}
			}
			C[iKey][types.State(jKey)] = sum
		}
	}

	return C
}

// MatPowHO computes P^n for HigherOrderMatrix using fast exponentiation.
func MatPowHO(P HigherOrderMatrix, n int, order []HigherOrderState) HigherOrderMatrix {
	if n == 0 {
		return HigherOrderMatrixIdentity(order)
	}
	res := HigherOrderMatrixIdentity(order)
	base := P
	for n > 0 {
		if n&1 == 1 {
			res = MatMulHO(res, base, order)
		}
		base = MatMulHO(base, base, order)
		n >>= 1
	}
	return res
}

// GenerateHigherOrderStates generates all possible histories of length K
// from the given set of states.
func GenerateHigherOrderStates(states []types.State, K int) []HigherOrderState {
	if K == 0 {
		return []HigherOrderState{{States: []types.State{}}}
	}

	// Recursively generate histories of length K-1
	subHistories := GenerateHigherOrderStates(states, K-1)
	var histories []HigherOrderState

	// Append each state to the histories of length K-1
	for _, subHistory := range subHistories {
		for _, state := range states {
			newHistory := HigherOrderState{
				States: append(append([]types.State{}, subHistory.States...), state),
			}
			histories = append(histories, newHistory)
		}
	}

	return histories
}
