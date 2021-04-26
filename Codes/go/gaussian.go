package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	p := plot.New()
	p.Title.Text = ""
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	gauss := plotter.NewFunction(func(x float64) float64 { return math.Exp(-x*x) })
	gauss.Samples = 100
	gauss.Color = color.RGBA{B: 255, A: 255}

	p.Add(gauss)

	p.X.Min = -3
	p.X.Max = 3
	p.Y.Min = 0
	p.Y.Max = 1.5

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "gaussian.png"); err != nil {
		panic(err)
	}
}