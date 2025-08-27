package stateMachines

import (
	"fmt"
	"markov_chains/types"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type CoinFlipChain struct {
	TransitionMatrix map[types.CoinFlipState]map[types.CoinFlipState]float64
	CurrentState     types.CoinFlipState
	StateCounter     map[types.CoinFlipState]int
	StateLog         []types.CoinFlipState
	LongestStreak    int
	currentStreak    int
	FirstPassage     int
	totalSteps       int
}

func (c *CoinFlipChain) RunSimulation(steps int) map[types.CoinFlipState]int {
	c.Reset()
	stateCounts := map[types.CoinFlipState]int{
		types.Heads: 0,
		types.Tails: 0,
	}
	for i := 0; i < steps; i++ {
		c.Step()
		stateCounts[c.CurrentState]++
	}
	return stateCounts
}

func (c *CoinFlipChain) PrintResults(name string, counts map[types.CoinFlipState]int, steps int) {
	fmt.Printf("=== %s Results ===\n", name)
	fmt.Printf("Total Steps: %d\n", steps)
	fmt.Printf("Heads: %d (%.2f%%)\n", counts[types.Heads], float64(counts[types.Heads])*100/float64(steps))
	fmt.Printf("Tails: %d (%.2f%%)\n", counts[types.Tails], float64(counts[types.Tails])*100/float64(steps))
	fmt.Printf("Longest Streak: %d\n", c.LongestStreak)
	fmt.Printf("First Passage Time to Change State: %d\n", c.FirstPassage)
	fmt.Printf("Entropy: %.4f\n", c.calculateEntropy())
	fmt.Println()
}

func (c *CoinFlipChain) Reset() {
	c.StateCounter = map[types.CoinFlipState]int{}
	c.CurrentState = types.Start
	c.StateLog = []types.CoinFlipState{}
	c.LongestStreak = 0
	c.totalSteps = 0
	c.FirstPassage = 0
}

func (c *CoinFlipChain) Step() {
	c.totalSteps = c.totalSteps + 1
	randVal := rand.Float64()
	nextState := c.CurrentState
	for state, prob := range c.TransitionMatrix[c.CurrentState] {
		if randVal < prob {
			nextState = state
			break
		}
		randVal -= prob
	}
	// Update streaks
	if nextState == c.CurrentState && nextState != types.Start {
		c.currentStreak++
	} else if nextState != types.Start {
		// We are switching.
		if c.FirstPassage == 0 && c.CurrentState != types.Start {
			c.FirstPassage = c.totalSteps
		}
		c.currentStreak = 1
	}
	if c.currentStreak > c.LongestStreak {
		c.LongestStreak = c.currentStreak
	}

	c.CurrentState = nextState
	c.StateCounter[nextState]++
	c.StateLog = append(c.StateLog, nextState)
}

// Add this method to CoinFlipChain
func (c *CoinFlipChain) PlotStateSequence(filename string) error {
	pts := make(plotter.XYs, len(c.StateLog))
	for i, state := range c.StateLog {
		var y float64
		switch state {
		case types.Heads:
			y = 1
		case types.Tails:
			y = 0
		default:
			y = -1
		}
		pts[i].X = float64(i)
		pts[i].Y = y
	}

	p := plot.New()
	p.Title.Text = "Coin Flip State Transitions"
	p.X.Label.Text = "Step"
	p.Y.Label.Text = "State"
	p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
		{Value: 0, Label: "Tails"},
		{Value: 1, Label: "Heads"},
	})

	line, err := plotter.NewLine(pts)
	if err != nil {
		return err
	}
	p.Add(line)

	return p.Save(8*vg.Inch, 3*vg.Inch, filename)
}

func (c *CoinFlipChain) calculateEntropy() float64 {
	pHeads := float64(c.StateCounter[types.Heads]) / float64(c.totalSteps)
	pTails := float64(c.StateCounter[types.Tails]) / float64(c.totalSteps)

	entropy := 0.0
	if pHeads > 0 {
		entropy -= pHeads * math.Log2(pHeads)
	}
	if pTails > 0 {
		entropy -= pTails * math.Log2(pTails)
	}
	return entropy
}
