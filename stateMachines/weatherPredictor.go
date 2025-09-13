package stateMachines

import (
	"markov_chains/helpers"
	"markov_chains/types"
)

type WeatherPredictor struct {
	*HigherOrderMatrixChain
}

func (wp *WeatherPredictor) RunSimulation(steps int) {
	wp.HigherOrderMatrixChain.RunSimulation(steps)
}

func NewArizonaWeatherPredictor() *WeatherPredictor {
	matrix := helpers.HigherOrderMatrix{
		helpers.HigherOrderState{States: []types.State{types.Weather_Sunny}}.StateKey():  {types.Weather_Sunny: 0.9, types.Weather_Rainy: 0.00, types.Weather_Cloudy: 0.1},
		helpers.HigherOrderState{States: []types.State{types.Weather_Rainy}}.StateKey():  {types.Weather_Sunny: 0.6, types.Weather_Rainy: 0.1, types.Weather_Cloudy: 0.3},
		helpers.HigherOrderState{States: []types.State{types.Weather_Cloudy}}.StateKey(): {types.Weather_Sunny: 0.9, types.Weather_Rainy: 0.1, types.Weather_Cloudy: 0.1},
	}
	return &WeatherPredictor{
		HigherOrderMatrixChain: NewHigherOrderMatrixChain(
			matrix,
			helpers.HigherOrderState{States: []types.State{types.Weather_Sunny}},
			"Weather Prediction Model for Arizona",
		),
	}
}

func NewPureRandomnessPredictor() *WeatherPredictor {
	possibleStates := []types.State{types.Weather_Cloudy, types.Weather_Rainy, types.Weather_Sunny, types.Weather_Snowy}

	// Iterate over possibleStates and assign 1/n possiblity to switch to all states.
	emptyMatrix := make(helpers.HigherOrderMatrix)
	numStates := len(possibleStates)
	for _, fromState := range possibleStates {
		emptyMatrix[helpers.HigherOrderState{States: []types.State{fromState}}.StateKey()] = make(map[types.State]float64)
		for _, toState := range possibleStates {
			emptyMatrix[helpers.HigherOrderState{States: []types.State{fromState}}.StateKey()][toState] = 1.0 / float64(numStates)
		}
	}
	return &WeatherPredictor{
		HigherOrderMatrixChain: NewHigherOrderMatrixChain(
			emptyMatrix,
			helpers.HigherOrderState{States: []types.State{types.Weather_Sunny}},
			"Weather Prediction Model for the most random place in the world",
		),
	}
}
