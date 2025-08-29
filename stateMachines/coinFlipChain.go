package stateMachines

import (
	"fmt"
	"image/color"
	"markov_chains/helpers"
	"markov_chains/types"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type CoinFlipChain struct { // Have this extend MatrixChain
	*MatrixChain // embedded pointer to MatrixChain
}

func (m *CoinFlipChain) RunSimulation(steps int) {
	m.runSimulation(steps)

	title := fmt.Sprintf("%s - State Probabilities", m.Description[0:20])
	err := m.plotStateProbabilitiesOverTime(title+".png", 100)

	if err != nil {
		fmt.Println("Error plotting:", err)
	}
	fmt.Printf("Output state probabilities line graph to %s\n", title+".png")
	fmt.Println()
}

func (m *CoinFlipChain) PlotStateSequence(filename string) error {
	pts := make(plotter.XYs, len(m.stateLog))
	for i, state := range m.stateLog {
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

func (m *CoinFlipChain) plotStateProbabilitiesOverTime(filename string, maxSteps int) error {
	headsProb := make(plotter.XYs, maxSteps+1)
	tailsProb := make(plotter.XYs, maxSteps+1)

	// Initial distribution: 100% Heads, 0% Tails
	dist := map[types.State]float64{
		types.CoinFlip_Heads: 0.50,
		types.CoinFlip_Tails: 0.50,
	}

	probabilityArray := helpers.GenerateStateProbabilityArray(maxSteps, m.TransitionMatrix, dist, []types.State{types.CoinFlip_Heads, types.CoinFlip_Tails})

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
