package stateMachines

import (
	"markov_chains/types"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type CoinFlipHigherOrderChain struct {
	*HigherOrderMatrixChain
}

func (m *CoinFlipHigherOrderChain) PlotStateSequence(filename string) error {
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
