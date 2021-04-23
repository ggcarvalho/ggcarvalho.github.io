package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

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
    for _,p := range S {
        sum_val += rectifier(p[len(p) - 1] - K)
    }
    C0 := math.Exp(-r*T)*sum_val / float64(I)

    duration := time.Since(start)
    fmt.Printf("European Option Value: %.3f\n", C0)
    fmt.Println("Execution time: ", duration)
}

// calculates max(x, 0)
func rectifier(x float64) float64 {
    if x >= 0.0 {
        return x
    }
    return 0.0
}
