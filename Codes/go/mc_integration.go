package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func monte_carlo_integral(function func(float64) float64, a float64, b float64, n int) float64 {
	s := 0.0
	for i := 0; i < n; i++ {
		u_i := rand.Float64()
		x_i := a + (b - a)*u_i
		s += function(x_i)
	}

	s = ( (b - a) / float64(n) ) * s
	return s
}

func gaussian(x float64) float64 {
	return math.Exp(-x*x)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	trials := 1000000
	fmt.Printf("Estimating the integral of f with %d point(s).\n\n", trials)

	integral := monte_carlo_integral(gaussian, -20.0, 20.0, trials)
	fmt.Printf("Approx. integral: %9f \n", integral)
}