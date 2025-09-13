package helpers

import (
	"markov_chains/types"
)

type Matrix = map[types.State]map[types.State]float64

// Identity matrix
func Identity(order []types.State) Matrix {
	I := make(Matrix, len(order))
	for _, s := range order {
		I[s] = make(map[types.State]float64, len(order))
		I[s][s] = 1.0
	}
	return I
}

func GenerateStateProbabilityArray(steps int, initialMatrix Matrix, initialDistribution map[types.State]float64, order []types.State) []map[types.State]float64 {
	probabilityArray := make([]map[types.State]float64, steps+1)
	probabilityArray[0] = initialDistribution

	for i := 1; i <= steps; i++ {
		probabilityArray[i] = make(map[types.State]float64)
		for _, state := range order {
			probabilityArray[i][state] = 0.0
			for prevState, prob := range probabilityArray[i-1] {
				probabilityArray[i][state] += prob * initialMatrix[prevState][state]
			}
		}
	}

	return probabilityArray
}

// Multiply two matrices
func MatMul(A, B Matrix, order []types.State) Matrix {
	C := make(Matrix, len(order))
	for _, i := range order {
		C[i] = make(map[types.State]float64, len(order))
		for _, j := range order {
			sum := 0.0
			for _, k := range order {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C
}

// Fast exponentiation: P^n
func MatPow(P Matrix, n int, order []types.State) Matrix {
	if n == 0 {
		return Identity(order)
	}
	res := Identity(order)
	base := P
	for n > 0 {
		if n&1 == 1 {
			res = MatMul(res, base, order)
		}
		base = MatMul(base, base, order)
		n >>= 1
	}
	return res
}
