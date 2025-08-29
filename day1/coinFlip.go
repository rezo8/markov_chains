package main

import "markov_chains/stateMachines"

var coinFlipChain = stateMachines.NewCoinFlipChain()
var evenlyBiasedCoinFlipChain = stateMachines.NewEvenlyBiasedCoinFlipChain()
var unEvenlyBiasedCoinFlipChain = stateMachines.NewUnEvenlyBiasedCoinFlipChain()

func main() {
	const steps = 10000
	coinFlipChain.RunSimulation(steps)
	evenlyBiasedCoinFlipChain.RunSimulation(steps)
	unEvenlyBiasedCoinFlipChain.RunSimulation(steps)

	const visuallyFriendlySteps = 500
	coinFlipChain.RunSimulation(visuallyFriendlySteps)
	evenlyBiasedCoinFlipChain.RunSimulation(visuallyFriendlySteps)
	unEvenlyBiasedCoinFlipChain.RunSimulation(visuallyFriendlySteps)

	coinFlipChain.PlotStateSequence("fair_coin_flip_chain.png")
	evenlyBiasedCoinFlipChain.PlotStateSequence("evenly_biased_coin_flip_chain.png")
	unEvenlyBiasedCoinFlipChain.PlotStateSequence("unevenly_biased_coin_flip_chain.png")

}
