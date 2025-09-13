package experiments

import (
	"markov_chains/stateMachines"
)

var arizonaWeatherPredictor = stateMachines.NewArizonaWeatherPredictor()
var randomWeatherSimulation = stateMachines.NewPureRandomnessPredictor()
var staySamePredictor = stateMachines.NewStayTheSamePredictor()

func RunWeatherPredictions() {
	const steps = 100000
	arizonaWeatherPredictor.RunSimulation(steps)
	randomWeatherSimulation.RunSimulation(steps)
	staySamePredictor.RunSimulation(steps)
}
