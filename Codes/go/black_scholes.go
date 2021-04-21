// import math
// from numpy import *
// from time import time
// # star import for shorter code
// random.seed(20000)
// t0 = time()
// # Parameters
// S0 = 100.; K = 105.; T = 1.0; r = 0.05; sigma = 0.2
// M = 50; dt = T / M; I = 250000
// # Simulating I paths with M time steps
// S = S0 * exp(cumsum((r - 0.5 * sigma ** 2) * dt
// + sigma * math.sqrt(dt)
// * random.standard_normal((M + 1, I)), axis=0))
// # sum instead of cumsum would also do
// # if only the final values are of interest
// S[0] = S0
// # Calculating the Monte Carlo estimator
// C0 = math.exp(-r * T) * sum(maximum(S[-1] - K, 0)) / I
// # Results output
// tnp2 = time() - t0

// print('The European Option Value is: ', C0)  # The European Option Value is:  8.165807966259603
// print('The Execution Time is: ',tnp2)  # The Execution Time is:  0.9024488925933838