package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const Pi float64 = math.Pi

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	numPoints := 1_000_000
	fmt.Printf("Estimating pi with %d point(s).\n\n", numPoints)

	sucess := 0
	for i := 0; i < numPoints; i++ {
		p := genRandomPoint()
		if isInsideCircle(p[0], p[1]) {sucess++}
	}

	piApprox := 4.0*(float64(sucess) / float64(numPoints))
	errorPct := 100.0*math.Abs(piApprox - Pi) / Pi

	fmt.Printf("Estimated pi: %9f \n", piApprox)
	fmt.Printf("pi: %9f \n", Pi)
	fmt.Printf("Error: %9f%%\n", errorPct)
}

// generates a random point p = (px, py)
func genRandomPoint() [2]float64 {
	px := 2.0*rand.Float64() - 1.0
	py := 2.0*rand.Float64() - 1.0
	return [2]float64{px, py}
}

// Condition to lie within the circular region
func isInsideCircle(x float64, y float64) bool {
	if x*x + y*y < 1 {
		return true
	}
	return false
}