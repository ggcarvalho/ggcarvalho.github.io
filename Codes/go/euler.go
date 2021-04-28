package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	numExperiments := 1_000_000
	fmt.Printf("Estimating e with %d experiment(s).\n\n", numExperiments)

	acc := 0
	for i := 0; i < numExperiments; i++ {
		sum := 0.0
		num2Sucess := 0

		for sum <= 1 {
			n := rand.Float64()
			sum += n
			num2Sucess++
		}
		acc += num2Sucess
	}

	expected := float64(acc) / float64(numExperiments)
	E := math.Exp(1)
	error_pct := 100.0 * math.Abs(expected-E) / E

	fmt.Printf("Expected vale: %9f \n", expected)
	fmt.Printf("e: %9f \n", E)
	fmt.Printf("Error: %9f%%\n", error_pct)
}
