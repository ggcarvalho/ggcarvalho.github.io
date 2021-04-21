package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func abs(x float64) float64 {
	if x < 0.0 {
		return -x
	}
	return x
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	trials := 10000000
	fmt.Printf("Calculating e with %d trial(s).\n\n", trials)

	acc := 0.0
	for i := 0; i < trials; i++ {
		sum := 0.0
		num2sucess := 0

		for sum <= 1 {
			n := rand.Float64()
			sum += n
			num2sucess += 1
		}
		acc += float64(num2sucess)
	}

	expected := acc / float64(trials)
	e := math.Exp(1)

	fmt.Printf("Expected vale: %9f \n", expected)
	fmt.Printf("e: %9f \n", e)

	error_pct := 100*abs(expected - e) / e

	fmt.Printf("Error: %9f%%\n", error_pct)

}