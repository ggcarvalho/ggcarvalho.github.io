package main

import (
	"image/color"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {

	p := plot.New()
	p.Title.Text = ""
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	upper := plotter.NewFunction(func(x float64) float64 { return math.Sqrt(1 - x*x) })
	upper.Samples = 500
	upper.Color = color.RGBA{B: 255, A: 255}

	lower := plotter.NewFunction(func(x float64) float64 { return -math.Sqrt(1 - x*x) })
	lower.Samples = 500
	lower.Color = color.RGBA{B: 255, A: 255}

	scatterData := randomPoints(250)
	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}

	s.Shape = draw.CircleGlyph{}
	s.Color = color.RGBA{R: 0, G: 135, B: 200, A: 255}

	p.Add(upper, lower, s)
	p.X.Min = -1
	p.X.Max = 1
	p.Y.Min = -1
	p.Y.Max = 1

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "circle.png"); err != nil {
		panic(err)
	}
}

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		pts[i].X = 2*rand.Float64() - 1.0
		pts[i].Y = 2*rand.Float64() - 1.0
	}
	return pts
}
