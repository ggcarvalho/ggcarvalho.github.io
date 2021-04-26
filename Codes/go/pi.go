package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	trials := 1000000
	fmt.Printf("Estimating pi with %d point(s).\n\n", trials)

	sucess := 0
	for i := 0; i < trials; i++ {
		p := gen_random_point()
		if inside_circle(p[0], p[1]) {sucess++}
	}

	pi_approx := 4.0*(float64(sucess) / float64(trials))
	pi := math.Pi
	error_pct := 100.0*abs(pi_approx - pi) / pi

	fmt.Printf("Estimated pi: %9f \n", pi_approx)
	fmt.Printf("pi: %9f \n", pi)
	fmt.Printf("Error: %9f%%\n", error_pct)
}

// generate random point p = (px, py)
func gen_random_point() [2]float64 {
	px := 2.0*rand.Float64() - 1.0
	py := 2.0*rand.Float64() - 1.0
	return [2]float64{px, py}
}

// Condition to lie within the circular region
func inside_circle(x float64, y float64) bool {
	if x*x + y*y < 1 {
		return true
	}
	return false
}

// absolute value of x
func abs(x float64) float64 {
	if x < 0.0 {
		return -x
	}
	return x
}