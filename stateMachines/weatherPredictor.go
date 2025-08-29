package stateMachines

import (
	"markov_chains/helpers"
	"markov_chains/types"
)

type WeatherPredictor struct {
	*MatrixChain
}

func (wp *WeatherPredictor) RunSimulation(steps int) {
	wp.MatrixChain.runSimulation(steps)
}

func NewArizonaWeatherPredictor() *WeatherPredictor {
	return &WeatherPredictor{
		MatrixChain: NewMatrixChain(
			helpers.Matrix{
				types.Weather_Sunny:  {types.Weather_Sunny: 0.9, types.Weather_Rainy: 0.00, types.Weather_Cloudy: 0.1},
				types.Weather_Rainy:  {types.Weather_Sunny: 0.6, types.Weather_Rainy: 0.1, types.Weather_Cloudy: 0.3},
				types.Weather_Cloudy: {types.Weather_Sunny: 0.9, types.Weather_Rainy: 0.1, types.Weather_Cloudy: 0.1},
			},
			"Weather Prediction Model for Arizona",
			nil,
		),
	}
}

func NewPureRandomnessPredictor() *WeatherPredictor {
	possibleStates := []types.State{types.Weather_Cloudy, types.Weather_Rainy, types.Weather_Sunny, types.Weather_Snowy}

	// Iterate over possibleStates and assign 1/n possiblity to switch to all states.
	emptyMatrix := make(helpers.Matrix)
	numStates := len(possibleStates)
	for _, fromState := range possibleStates {
		emptyMatrix[fromState] = make(map[types.State]float64)
		for _, toState := range possibleStates {
			emptyMatrix[fromState][toState] = 1.0 / float64(numStates)
		}
	}
	return &WeatherPredictor{
		MatrixChain: NewMatrixChain(
			emptyMatrix,
			"Weather Prediction Model for the most random place in the world",
			nil,
		),
	}
}
