package main

import (
	"fmt"
	"time"
	"image/color"
	"math"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func rectifier(x float64) float64 {
	if x >= 0.0 {
		return x
	}
	return 0.0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()

	// Parameters
	S0 := 100.0 // initial value
	K := 105.0 // strike price
	T := 1.0 // maturity
	r := 0.05 //risk free short rate
	sigma := 0.2 //volatility
	M := 50 // number of time steps
	dt := T / float64(M) //length of time interval
	I := 250000 // number of paths

	var S [][]float64

	// Simulating I paths with M time steps
	for i := 1; i < I; i++ {
		var path []float64
		for t := 0; t <= M; t++ {
			if t == 0 {
				path = append(path, S0)
			} else {
				z := rand.NormFloat64()
				St := path[t - 1]*math.Exp((r - 0.5*(sigma*sigma))*dt + sigma*math.Sqrt(dt)*z)
				path = append(path, St)
			}
		}
		S = append(S, path)
	}

	// Calculating the Monte Carlo estimator
	sum_val := 0.0
	var end_inner []float64
	var end []float64
	for _,p := range S {
		sum_val += rectifier(p[len(p) - 1] - K)
		end_inner = append(end_inner, rectifier(p[len(p) - 1] - K))
		end = append(end, p[len(p) - 1])
	}
	C0 := math.Exp(-r*T)*sum_val / float64(I)

	duration := time.Since(start)
	fmt.Printf("European Option Value: %.3f\n", C0)
	fmt.Println("Execution time: ", duration)

	// Plots
	//Histogram of all simulated end-of-period index level values
	histPlot(end, 50, "", "index level", "frequency","end_hist")

	// Histogram of all simulated end-of-period option inner values
	histPlot(end_inner, 50, "", "option inner value", "frequency", "end_inner_hist")

	var paths [][]float64
	for i := 0; i < 50; i++ {
		paths = append(paths, S[i])
	}

	p := plot.New()

	for _, pth := range paths {
		current_path := points(pth)
		p.Title.Text = ""
		p.X.Label.Text = "time step"
		p.Y.Label.Text = "index level"

		// Make a line plotter with points and set its style.
		lpLine, lpPoints, err := plotter.NewLinePoints(current_path)
		if err != nil {
			panic(err)
		}
		lpLine.Color = color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)),A: 255}
		lpPoints.Shape = draw.PyramidGlyph{}
		lpPoints.Color = color.RGBA{R: 255, A: 255}

		p.Add(lpLine)
	}

	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "paths.png"); err != nil {
		panic(err)
	}
}


func histPlot(values plotter.Values, bins int, title string, xLabel string, yLabel string, savename string) {
    p := plot.New()
    p.Title.Text = title
	p.X.Label.Text = xLabel
    p.Y.Label.Text = yLabel
    hist, err := plotter.NewHist(values, bins)
    if err != nil {
        panic(err)
    }
    p.Add(hist)

    if err := p.Save(8*vg.Inch, 6*vg.Inch, savename + ".png"); err != nil {
        panic(err)
    }
}

func points(path []float64) plotter.XYs {
	pts := make(plotter.XYs, len(path))

	j := 0.0
	for i := range pts {

		pts[i].X = j
		pts[i].Y = path[i]
		j = j + 1.0
	}
	return pts
}