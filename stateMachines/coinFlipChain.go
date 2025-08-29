package stateMachines

import (
	"fmt"
	"image/color"
	"markov_chains/helpers"
	"markov_chains/types"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type CoinFlipChain struct {
	TransitionMatrix helpers.Matrix ///map[types.State]map[types.State]float64
	CurrentState     types.State
	StateCounter     map[types.State]int
	StateLog         []types.State
	Description      string
	LongestStreak    int
	currentStreak    int
	FirstPassage     int
	totalSteps       int
}

func (c *CoinFlipChain) RunSimulation(steps int) map[types.State]int {
	c.reset()
	stateCounts := map[types.State]int{
		types.CoinFlip_Heads: 0,
		types.CoinFlip_Tails: 0,
	}
	for i := 0; i < steps; i++ {
		c.step()
		stateCounts[c.CurrentState]++
	}

	c.printResults()
	return stateCounts
}

func (c *CoinFlipChain) PlotStateSequence(filename string) error {
	pts := make(plotter.XYs, len(c.StateLog))
	for i, state := range c.StateLog {
		var y float64
		switch state {
		case types.CoinFlip_Heads:
			y = 1
		case types.CoinFlip_Tails:
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

func (c *CoinFlipChain) PredictNthState(n int) helpers.Matrix {
	return helpers.MatPow(c.TransitionMatrix, n, []types.State{types.CoinFlip_Heads, types.CoinFlip_Tails})
}

func (c *CoinFlipChain) plotStateProbabilitiesOverTime(filename string, maxSteps int) error {
	headsProb := make(plotter.XYs, maxSteps+1)
	tailsProb := make(plotter.XYs, maxSteps+1)

	// Initial distribution: 100% Heads, 0% Tails
	dist := map[types.State]float64{
		types.CoinFlip_Heads: 0.50,
		types.CoinFlip_Tails: 0.50,
	}

	probabilityArray := helpers.GenerateStateProbabilityArray(maxSteps, c.TransitionMatrix, dist, []types.State{types.CoinFlip_Heads, types.CoinFlip_Tails})

	// iterate over and plot probabilityArray
	for n := 0; n <= maxSteps; n++ {
		headsProb[n].X = float64(n)
		headsProb[n].Y = probabilityArray[n][types.CoinFlip_Heads]
		tailsProb[n].X = float64(n)
		tailsProb[n].Y = probabilityArray[n][types.CoinFlip_Tails]
	}

	p := plot.New()
	p.Title.Text = "Probability of Heads/Tails at Step N"
	p.X.Label.Text = "Step N"
	lineHeads, err := plotter.NewLine(headsProb)
	if err != nil {
		return err
	}
	lineHeads.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Red for Heads
	lineHeads.LineStyle.Width = vg.Points(2)
	lineHeads.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	lineTails, err := plotter.NewLine(tailsProb)
	if err != nil {
		return err
	}
	lineTails.Color = color.RGBA{B: 255, A: 255} // Blue for Tails
	lineTails.LineStyle.Width = vg.Points(2)

	p.Add(lineHeads, lineTails)
	p.Legend.Add("Heads", lineHeads)
	p.Legend.Add("Tails", lineTails)
	p.Y.Min = 0
	p.Y.Max = 1

	return p.Save(8*vg.Inch, 4*vg.Inch, filename)
}

func (c *CoinFlipChain) printResults() {
	fmt.Printf("=== %s Results ===\n", c.Description)
	fmt.Printf("Total Steps: %d\n", len(c.StateLog))
	fmt.Printf("Heads: %d (%.2f%%)\n", c.StateCounter[types.CoinFlip_Heads], float64(c.StateCounter[types.CoinFlip_Heads])*100/float64(len(c.StateLog)))
	fmt.Printf("Tails: %d (%.2f%%)\n", c.StateCounter[types.CoinFlip_Tails], float64(c.StateCounter[types.CoinFlip_Tails])*100/float64(len(c.StateLog)))
	fmt.Printf("Longest Streak: %d\n", c.LongestStreak)
	fmt.Printf("First Passage Time to Change State: %d\n", c.FirstPassage)
	fmt.Printf("Entropy: %.4f\n", c.calculateEntropy())

	title := fmt.Sprintf("%s - State Probabilities", c.Description[0:20])
	err := c.plotStateProbabilitiesOverTime(title+".png", 100)

	if err != nil {
		fmt.Println("Error plotting:", err)
	}
	fmt.Printf("Output state probabilities line graph to %s\n", title+".png")
	fmt.Println()
}

func (c *CoinFlipChain) reset() {
	c.StateCounter = map[types.State]int{}
	c.CurrentState = types.CoinFlip_Start
	c.StateLog = []types.State{}
	c.LongestStreak = 0
	c.totalSteps = 0
	c.FirstPassage = 0
}

func (c *CoinFlipChain) step() {
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
	if nextState == c.CurrentState && nextState != types.CoinFlip_Start {
		c.currentStreak++
	} else if nextState != types.CoinFlip_Start {
		// We are switching so mark first passage.
		if c.FirstPassage == 0 && c.CurrentState != types.CoinFlip_Start {
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

func (c *CoinFlipChain) calculateEntropy() float64 {
	pHeads := float64(c.StateCounter[types.CoinFlip_Heads]) / float64(c.totalSteps)
	pTails := float64(c.StateCounter[types.CoinFlip_Tails]) / float64(c.totalSteps)

	entropy := 0.0
	if pHeads > 0 {
		entropy -= pHeads * math.Log2(pHeads)
	}
	if pTails > 0 {
		entropy -= pTails * math.Log2(pTails)
	}
	return entropy
}
