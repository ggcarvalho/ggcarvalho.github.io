---
title: "The art of solving problems with Monte Carlo simulations"
type: post
description: "Using the power of randomness to answer scientific questions."
date: "2021-05-03"
image: "/img/posts/montecarlo/dies.jpg"
tag: "go"
name: "go"
hashtags: "#go #golang #montecarlo #quantitativeresearch #trade #algotrade #quant #finance"
draft: false
---

In this article, we will explore some examples and applications of Monte Carlo simulations using the Go programming language. To keep this article fun and interactive, after each Go code provided you will find a link to the <em>Go Playground</em>, where you can run it without installing Go on your machine.

Put your adventure helmets on!

<div style="text-align:left"><img src="/img/posts/montecarlo/gopher_adventurer.jpg" style="width: 20%"></div>

{{< table_of_contents >}}

## Requirements

As stated previously, there is no need to install anything on your computer, you can use the Go Playground. However, if you wish to run the programs locally on your computer (which I recommend), you should <a href="https://golang.org/" target="_blank">download and install Go</a>. If you want to learn the Go Programming Language, check the "Recommended Reading" section at the end of this article, as well as

- <a href="https://gobyexample.com/" target="_blank">Go by Example</a>

- <a href="https://tour.golang.org/welcome/1" target="_blank">A Tour of Go</a>

- <a href="https://golang.org/doc/effective_go" target="_blank">Effective Go</a>

- <a href="https://golang.org/doc/code" target="_blank">How to Write Go Code</a>

The Go programming language is deemed to be <strong>the most promising programming language</strong> today due to its speed and simplicity, and I recommend you to at least get acquainted with it.

## Quick Introduction

---

Generally speaking, Monte Carlo methods (or simulations) consist of a broad class of computational algorithms that rely on repeated random sampling to obtain numerical results. This technique is used throughout areas such as physics, finance, engineering, project management, insurance, and transportation, where a numerical result is needed and the underlying theory is difficult and/or unavailable.

It was invented by <a href="https://en.wikipedia.org/wiki/John_von_Neumann" target="_blank">John von Neumann</a>, <a href="https://en.wikipedia.org/wiki/Stanislaw_Ulam" target="_blank">Stanisław Ulam</a>, and <a href="https://en.wikipedia.org/wiki/Nicholas_Metropolis" target="_blank">Nicholas Metropolis</a>, who were employed on a secret assignment in the Los Alamos National Laboratory, while working on a nuclear weapon project called the Manhattan Project. It was named after a well-known casino town, called Monaco since chance and randomness are core to the modeling approach, similar to a game of roulette.

Here are two excerpts taken from books on Monte Carlo simulations. The first comes from N.T. Thomopoulos' "Essentials of Monte Carlo Simulation: Statistical Methods for Building Simulation Models", and the second comes from Paul Glasserman's "Monte Carlo Methods in Financial Engineering (Stochastic Modelling and Applied Probability)":

>To apply the Monte Carlo method, the analyst constructs a mathematical model that simulates a real system. A large number of random sampling of the model is applied yielding a large number of random samples of output results from the model. [...] The method is based on running the model many times as in random sampling. For each sample, random variates are generated on each input variable; computations are run through the model yielding random outcomes on each output variable. Since each input is random, the outcomes are random. In the same way, they generated thousands of such samples and achieved thousands of outcomes for each output variable.

>Monte Carlo methods are based on the analogy between probability and volume. The mathematics of measure formalizes the intuitive notion of probability, associating an event with a set of outcomes and defining the probability of the event to be its volume or measure relative to that of a universe of possible outcomes. Monte Carlo uses this identity in reverse, calculating the volume of a set by interpreting the volume as a probability. In the simplest case, this means sampling randomly from a universe of possible outcomes and taking the fraction of random draws that fall in a given set as an estimate of the set’s volume. The law of large numbers ensures that this estimate converges to the correct value as the number of draws increases. The central limit theorem provides information about the likely magnitude of the error in the estimate after a finite number of draws.

If you are new to the subject, keep these ideas in mind as they will help you to understand what follows next.

## First Examples

---

Let's start our journey with some basic, and even textbook, examples of Monte Carlo simulations. You're going to notice a pattern when implementing this method. Use it when implementing your simulations.

### Estimating $\pi$

This is by far the most famous example of Monte Carlo simulations, considered to be the "zeroth example" of the subject. To estimate $\pi$ we need to pose the problem in probabilistic terms. We do so by considering a circle of radius $r=1$ inscribed in a square of side $l=2$, both centered in the origin of a cartesian coordinate system. This situation is depicted below.

<div style="text-align:center"><img src="/img/posts/montecarlo/circle.png" style="width: 30%"></div>

If we were to draw random points in this square, some will fall within the circle and some won't. But the ratio between points inside the circular region and the total amount of points we draw will be closer and closer to the ratio between the area of the circle and the area of the square as we draw more of these random points.

<div style="text-align:center"><img src="/img/posts/montecarlo/pi_simulation.gif" style="width: 30%"></div>

As you might know, the area of the circular region is $A_{\bigcirc} = \pi\cdot r^2$ and the area of the square is $A_{\square} = (2r)^2 = 4r^2$. Thus, $$\pi = 4\cdot \frac{A_{\bigcirc}}{A_{\square}}.$$

As a result, we can estimate $\pi$ as
$$\pi \approx 4\cdot \frac{\text{$\\#$ points inside the circle}}{\text{$\\#$ points}}.$$
Remember that, in our case, $(x, y)$ will fall inside the circular region if $x^2 + y^2 < 1$.

Now, let's turn this idea into a Go code.

```go
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
        if isInsideCircle(p[0], p[1]) {
            sucess++
        }
    }

    piApprox := 4.0 * (float64(sucess) / float64(numPoints))
    errorPct := 100.0 * math.Abs(piApprox-Pi) / Pi

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
    if x*x+y*y < 1 {
        return true
    }
    return false
}
```

<a href="https://play.golang.org/p/bPWYH3lxm_S" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run pi.go

Estimating pi with 1000000 point(s).

Estimated pi:  3.142864
pi:  3.141593
Error:  0.040468%
```

We were able to approximate $\pi$ with an error of $0.040468$% ! Check this video to see how many digits of $\pi$ you really need.

<div class="container" style="width:50%">
  <div class="center">
    {{<youtube l-vHGf4j90Y>}}
  </div>
</div>

Let's jump to the next example.

### Estimating Euler's Number

Not long ago, Lex Fridman (AI researcher, YouTuber, and Podcast Host) published the following in a LinkedIn post:

<div style="text-align:center"><img src="/img/posts/montecarlo/lex.png" style="width: 60%"></div>

This is very intriguing! Since I have no idea how to prove this statement, I decided to write a simple Python program to test it myself. I wrote my code as a response to Lex's post, and people were surprised by it (this is the actual motivation to write this blog post).
This, once again, showed me the power of Monte Carlo simulations in scientific computing, where you can find precise numerical answers to your problems without relying on any theoretical background (they might not even exist!).

I wrote an equivalent Go code to solve this problem. Here it is:

```go
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
```

<a href="https://play.golang.org/p/XH0wYDYHHcP" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run euler.go

Estimating e with 1000000 experiment(s).

Expected vale:  2.718631
e:  2.718282
Error:  0.012845%
```

An astonishing result!

### The Birthday Paradox

This is a famous problem in statistics:
>In a group of $23$ people, the probability of a shared birthday exceeds $50$%.

That sounds weird at first, but if you're good enough with math you can easily prove this statement. However, we are not interested in formal proofs here. That is the whole point of these simulations. The idea is very simple: create a list with $n$ random numbers (in our case $n = 23$) between $0$ and $364$ representing the birth day of each person and if (at least) two of them coincide, you increment the `success` variable. Do it a certain number of times and calculate the probability dividing the number of successes by the total number of simulations.

The corresponding Go code:

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    numPeople := 23
    trials := 1_000_000
    sucess := 0
    for i := 0; i < trials; i++ {
        bdays := genBdayList(numPeople)
        uniques := uniqueSlice(bdays)

        if !(len(bdays) == len(uniques)) {
            sucess++
        }
    }
    probability := float64(sucess) / float64(trials)
    fmt.Printf("The probability of at least 2 persons in a group of %d people share a birthday is %.2f%%\n", numPeople, 100.0*probability)
}

// returns a slice with the uniqueSlice elements of a given slice
func uniqueSlice(s []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range s {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

// generates the list of birth days
func genBdayList(n int) []int {
    var bdays []int
    for i := 0; i < n; i++ {
        bday := rand.Intn(365)
        bdays = append(bdays, bday)
    }
    return bdays
}
```

<a href="https://play.golang.org/p/SdX37thwMYK" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run birthday.go

The probability of at least 2 persons in a group of 23 people share a birthday is 50.67%
```

You can modify this code to reproduce the following table from <a href="https://en.wikipedia.org/wiki/Birthday_problem" target="_blank">Wikipedia</a>:

<table class="table table-striped">
  <thead>
    <tr>
      <th scope="col">$n$</th>
      <th scope="col">$\mathcal{P}(n)$</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th scope="row">$1$</th>
      <td>$0.0$%</td>
    </tr>
    <tr>
      <th scope="row">$5$</th>
      <td>$2.7$%</td>
    </tr>
    <tr>
      <th scope="row">$10$</th>
      <td>$11.7$%</td>
    </tr>
    <tr>
      <th scope="row">$20$</th>
      <td>$41.1$%</td>
    </tr>
    <tr>
      <th scope="row">$23$</th>
      <td>$50.7$%</td>
    </tr>
    <tr>
      <th scope="row">$30$</th>
      <td>$70.6$%</td>
    </tr>
    <tr>
      <th scope="row">$40$</th>
      <td>$89.1%$%</td>
    </tr>
    <tr>
      <th scope="row">$50$</th>
      <td>$97.0%$%</td>
    </tr>
    <tr>
      <th scope="row">$60$</th>
      <td>$99.4%$%</td>
    </tr>
    <tr>
      <th scope="row">$70$</th>
      <td>$99.9%$%</td>
    </tr>
    <tr>
      <th scope="row">$75$</th>
      <td>$99.97%$%</td>
    </tr>
    <tr>
      <th scope="row">$100$</th>
      <td>$99.99997$%</td>
    </tr>
    <tr>
      <th scope="row">$\geq 365$</th>
      <td>$100$%</td>
    </tr>
</table>

Of course, for $n \geq 365$ you don't need any calculations, it's a straightforward consequence of the <a href="https://en.wikipedia.org/wiki/Pigeonhole_principle" target="_blank"><em>pigeonhole principle</em></a>.

## The (In)Famous Monty Hall Problem

---

This is a problem that has been confusing people for ages. Just like the birthday problem, you can solve it using basic math/probability theory, which we won't do. Let's state the problem and provide a Monte Carlo simulation to solve it.

Here is the problem:
>Suppose you're on a game show, and you're given the choice of three doors: Behind one door is a valuable prize; behind the others, goats. You pick a door, say No. $1$, and the host, who knows what's behind the doors, opens another door, say No. $3$, which has a goat. He then says to you, "Do you want to pick door No. $2$?" Is it to your advantage to switch your choice?

<div style="text-align:center"><img src="/img/posts/montecarlo/montyhall.png" style="width: 60%; margin: 2%"></div>

People often believe that they are in a $50-50$ situation and therefore the switch is not very relevant. But if you carefully solve this problem, you will find that there is a probability of $2/3 \approx 66.7$% to win the prize if you decide to switch doors.

The following Go code simulates this game and estimates the probability of winning if the guest chooses to switch doors. Note that we first set the game so that the doors are properly chosen at random, and then we simulate several games to estimate the desired probability.

```go
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    numGames := 10_000_000
    fmt.Printf("Estimating the probability of winning by switching doors with %d game(s).\n\n", numGames)

    sucess := 0
    for i := 0; i < numGames; i++ {
        newDoor, prizeDoor := setMontyHallGame()
        if newDoor == prizeDoor {
            sucess++
        }
    }
    probability := float64(sucess) / float64(numGames)
    theoreticalValue := 2.0 / 3.0

    errorPct := 100.0 * math.Abs(probability-theoreticalValue) / theoreticalValue

    fmt.Printf("Estimated probability: %9f \n", probability)
    fmt.Printf("Theoretical value: %9f \n", theoreticalValue)
    fmt.Printf("Error: %9f%%\n", errorPct)
}

// randomly sets the game
func setMontyHallGame() (int, int) {
    var montysChoice int
    var prizeDoor int
    var goat1Door int
    var goat2Door int
    var newDoor int

    guestDoor := rand.Intn(3)

    areDoorsSelected := false
    for !areDoorsSelected {
        prizeDoor = rand.Intn(3)
        goat1Door = rand.Intn(3)
        goat2Door = rand.Intn(3)
        if prizeDoor != goat1Door && prizeDoor != goat2Door && goat1Door != goat2Door {
            areDoorsSelected = true
        }
    }

    showGoat := false
    for !showGoat {
        montysChoice = rand.Intn(3)
        if montysChoice != prizeDoor && montysChoice != guestDoor {
            showGoat = true
        }
    }

    madeSwitch := false
    for !madeSwitch {
        newDoor = rand.Intn(3)
        if newDoor != guestDoor && newDoor != montysChoice {
            madeSwitch = true
        }
    }
    return newDoor, prizeDoor
}
```

<a href="https://play.golang.org/p/1j5oSdmBJjt" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run monty_hall.go

Estimating the probability of winning by switching doors with 10000000 game(s).

Estimated probability:  0.666447
Theoretical value:  0.666667
Error:  0.032950%
```

Therefore, contrary to popular belief, it is more advantageous to the guest to switch doors, confirming the theoretical result.

## Integration Using Monte Carlo Simulations

---

Now, let's see how we can use the Monte Carlo method to find the value of definite integrals of continuous functions in a specified range. This method is particularly useful for higher-dimensional integrals, due to its convergence properties.

Just as a reminder, if $f: [a,b] \rightarrow \mathbb{R}$ is a continuous function, then the quantity $$S = \int_a^b f(x)dx,$$
represents the area of the region between the graph of $f$ and the $x-$axis.

<div style="text-align:center"><img src="/img/posts/montecarlo/integralarea.png" style="width: 30%; margin: 2%"></div>

One important feature of this property is that this area has a sign, having negative values if the region is below the $x-$axis.

There are some ways to approximate this area, such as Newton-Cotes rules, trapezoidal rule, and Simpson's rule. However, one clever way to numerically integrate continuous functions is using the formula  $$S \approx \frac{b-a}{n}\sum_{i=1}^n f(a + (b-a)U_i),$$
where $U_i \sim \mathcal{U}(0,1)$, i.e. the $U_i$ are uniformly distributed in $[0,1]$ (feel free to try different probability distributions and compare the results). The Monte Carlo integration is

We are going to use this technique to solve a classic problem. If you are a calculus geek, you might know how difficult it is to calculate the integral $$S = \int_{-\infty}^{\infty} e^{-x^2}dx.$$

It involves a trick using Fubini's theorem and a change from cartesian to polar coordinates. Surprisingly, the result of this integral is $\sqrt{\pi}$. Let's use Monte Carlo integration to evaluate $$\bar{S} = \int_{-20}^{20} e^{-x^2}dx.$$

<div style="text-align:center"><img src="/img/posts/montecarlo/gaussian.png" style="width: 35%; margin: 2%"></div>

We see that $f$ rapidly decreases when moving away from $x=0$, so the definite integral in $[-20,20]$ seems to be a good approximation.

The corresponding Go code is:

```go
package main

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    numPoints := 1_000_000
    fmt.Printf("Estimating the integral of f with %d point(s).\n\n", numPoints)

    integral := monteCarloIntegrator(gaussian, -20.0, 20.0, numPoints)
    fmt.Printf("Approx. integral: %9f \n", integral)
}

// MC integrator
func monteCarloIntegrator(function func(float64) float64, a float64, b float64, n int) float64 {
    s := 0.0
    for i := 0; i < n; i++ {
        u_i := rand.Float64()
        x_i := a + (b - a)*u_i
        s += function(x_i)
    }

    s = ((b - a) / float64(n)) * s
    return s
}

// function to be integrated
func gaussian(x float64) float64 {
    return math.Exp(-x*x)
}
```

<a href="https://play.golang.org/p/JqduOVngdiS" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run mc_integration.go

Estimating the integral of f with 1000000 point(s).

Approx. integral:  1.771559
```

In fact, $1.771559^2 \approx 3.138$. You should use the Go Playground to experiment with different functions (quadratic, cubic, polynomials, or even more complicated functions).

## Option Pricing Using the Black-Scholes Model

---

For the final section of this article, I have something special that draws a lot of attention: the Black-Scholes model.

The Black–Scholes, or Black–Scholes–Merton model (<a href="https://en.wikipedia.org/wiki/Fischer_Black" target="_blank">Fischer Black</a>, <a href="https://en.wikipedia.org/wiki/Myron_Scholes" target="_blank">Myron Scholes</a>, and <a href="https://en.wikipedia.org/wiki/Robert_C._Merton" target="_blank">Robert C Merton</a>) , is a mathematical model for the dynamics of a financial market containing derivative investment instruments, giving a theoretical estimate of the price of <em>European-style options</em> and shows that the option has a unique price given the risk of the security and its expected return. This work granted Myron Scholes and Robert C Merton their Nobel Prize in Economics ($1997$), and has been widely used in algorithmic trading strategies around the world.

### The Equation

We start with the Black-Scholes-Merton formula ($1973$) for the pricing of European call options on an underlying (e.g stocks and indexes) without dividends:
\begin{eqnarray*}
C(S_t, K, t, T, r, \sigma) &=& S_t\cdot N(d_1) - e^{-r(T-t)}\cdot K \cdot N(d_2)\newline\newline
N(d) &=& \frac{1}{\sqrt{2\pi}}\int_{-\infty}^d e^{-\frac{1}{2}x^2}dx \newline\newline
d_1 &=& \frac{\log\frac{S_t}{K} + (T-t)\left(r + \frac{\sigma^2}{2}\right)}{\sigma\sqrt{T-t}}\newline\newline
d_2 &=& \frac{\log\frac{S_t}{K} + (T-t)\left(r - \frac{\sigma^2}{2}\right)}{\sigma\sqrt{T-t}}.
\end{eqnarray*}

In the equations above $S_t$ is the price of the underlying at time $t$, $\sigma$ is the constant volatility (standard deviation of returns) of the underlying, $K$ is the strike price of the option, $T$ is the maturity date of the option, $r$ is the risk-free short rate.

The Black-Scholes-Merton ($1973$) stochastic differential equation is given by $$dS_t = rS_t dt + \sigma S_t dZ_t,$$
where $Z(t)$ is the random component of the model (a Brownian motion). In this model, the risky underlying follows, under risk neutrality, a geometric Brownian motion with a stochastic differential equation (SDE).

We will look at the discretized version of the BSM model (Euler discretization), given by $$S_t = S_{t-\Delta t}  +  \exp\left(\left(r - \frac{\sigma^2}{2}\right)\Delta t + \sigma\sqrt{\Delta t}z_t \right).$$

The variable $z$ is a standard normally distributed random variable, $0 < \Delta t < T$, a (small
enough) time interval. It also holds $0 < t \leq T$ with $T$ the final time horizon.

In this simulation we use the values $S_0 = 100$, $K = 105$, $T = 1.0$, $r = 0.05$, $\sigma = 0.2$. Let's see what is the expected option price using these parameters and assuming $t=0$, then we will run a Monte Carlo simulation to find the option price under the same conditions.

### Option Pricing

We are going to use the first set of equations, together with our Monte Carlo integrator, to calculate the option price under the conditions established. The corresponding Go code is:

```go
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())

    // Parameters
    S0 := 100.0  // initial value
    K := 105.0   // strike price
    T := 1.0     // maturity
    r := 0.05    //risk free short rate
    sigma := 0.2 //volatility
    numPoints := 250_000

    start := time.Now()
    optionPrice := bsmCallValue(S0, K, T, r, sigma, numPoints)
    duration := time.Since(start)

    fmt.Printf("European Option Value: %.3f\n", optionPrice)
    fmt.Println("Execution time: ", duration)
}

func bsmCallValue(S0 float64, K float64, T float64, r float64, sigma float64, n int) float64 {
    d1 := math.Log(S0/K) + T*(r+0.5*sigma*sigma)/(sigma*math.Sqrt(T))
    d2 := math.Log(S0/K) + T*(r-0.5*sigma*sigma)/(sigma*math.Sqrt(T))

    value := S0*monteCarloIntegrator(gaussian, -20.0, d1, n) - K*math.Exp(-r*T)*monteCarloIntegrator(gaussian, -20.0, d2, n)

    return value
}

// MC integrator
func monteCarloIntegrator(function func(float64) float64, a float64, b float64, n int) float64 {
    s := 0.0
    for i := 0; i < n; i++ {
        u_i := rand.Float64()
        x_i := a + (b-a)*u_i
        s += function(x_i)
    }
    s = ((b - a) / float64(n)) * s
    return s
}

// function to be integrated
func gaussian(x float64) float64 {
    return (1 / math.Sqrt(2*math.Pi)) * math.Exp(-0.5*x*x)
}
```

<a href="https://play.golang.org/p/gUNpYZOvzJX" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run option_pricing.go

European Option Value: 7.964
Execution time:  12.171679ms
```

This is our benchmark value for the Monte Carlo estimator to follow.

### The Simulation

We follow the steps:

1. Divide the time interval $[0, T]$ in equidistant subintervals of length $\Delta t$.
2. Start iterating $i = 1, 2,..., I$.
    - For every time step $t  \in \\{\Delta t, 2\Delta t,..., T \\}$, draw pseudo-random numbers $z_t(i)$.
    - Determine the time $T$ value of the index level $S_T(i)$ by applying the pseudo-random numbers time step by time step to the discretized equation.
    - Determine the inner value $h_T$ of the European call option at $T$ as $h_T(S_T(i)) = \max(S_T(i) – K, 0)$.
    - Iterate until $i = I$.

3. Sum up the inner values, average, and discount them back with the riskless short rate according to the formula
$$C_0 \approx e^{-rT} \frac{1}{I} \sum_I h_T(S_T(i)),$$ called the Monte Carlo estimator for the European call option.

Without any further ado, here is the corresponding Go code:

```go
package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    start := time.Now()

    // Parameters
    S0 := 100.0          // initial value
    K := 105.0           // strike price
    T := 1.0             // maturity
    r := 0.05            //risk free short rate
    sigma := 0.2         //volatility
    M := 50              // number of time steps
    dt := T / float64(M) //length of time interval
    I := 250_000         // number of paths/simulations
    var S [][]float64

    // Simulating numPaths paths with M time steps
    for i := 1; i < I; i++ {
        var path []float64
        for t := 0; t <= M; t++ {
            if t == 0 {
                path = append(path, S0)
            } else {
                z := rand.NormFloat64()
                St := path[t-1] * math.Exp((r-0.5*(sigma*sigma))*dt+sigma*math.Sqrt(dt)*z)
                path = append(path, St)
            }
        }
        S = append(S, path)
    }

    // Calculating the Monte Carlo estimator
    sumVal := 0.0
    for _, p := range S {
        sumVal += rectifier(p[len(p)-1] - K)
    }
    C0 := math.Exp(-r*T) * sumVal / float64(I)
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
```

<a href="https://play.golang.org/p/tcEbs3jZgNu" target="_blank">Run this code in the Go Playground</a>

```bash
$ go run black_scholes.go

European Option Value: 8.027
Execution time:  430.464289ms
```

We got a very satisfactory result using the Monte Carlo estimator (remember that the value was $7.964$ using the BSM formula and our Monte Carlo integrator).

Let's compare the results with the same simulation in Python (taken from Yves Hilpisch's "Python for Finance"):

<div style="text-align:center"><img src="/img/posts/montecarlo/yves.png" style="width: 50%; margin: 2%"></div>

Nearly the same result in a fraction of the time! To be completely fair, when the author uses full Numpy vectorization the results are much better in terms of performance, although we still have a clear winner.

<div style="text-align:center"><img src="/img/posts/montecarlo/yves2.png" style="width: 50%; margin: 2%"></div>

### Graphical Analysis

First, let's plot the simulated index levels (the paths taken during the simulation). The figures below represent the first $10$, the first $100$, and $250,000$ (total number of paths) simulated index levels respectively.

<div style= "text-align:center">
<a href="/img/posts/montecarlo/10.png" target="_blank"><img src="/img/posts/montecarlo/10.png"  alt="Mean" style="width:45%; margin:1%"></a>
<a href="/img/posts/montecarlo/100.png" target="_blank"><img src="/img/posts/montecarlo/100.png"  alt="Gaussian" style="width:45%; margin:1%"></a>
</div>
<div style= "text-align:center">
<a href="/img/posts/montecarlo/all.png" target="_blank"><img src="/img/posts/montecarlo/all.png"  alt="Sharpen" style="width:45%; margin:1%"></a>
</div>

You need to appreciate the fact that each path taken by the index level is an actual possible path, and the option price is calculated by taking every possibility into account.

Second, we want to see the frequency of the simulated index levels at the end of the simulation period.

<div style= "text-align:center">
<a href="/img/posts/montecarlo/end_hist.png" target="_blank"><img src="/img/posts/montecarlo/end_hist.png"  alt="Sharpen" style="width:50%; margin:1%"></a>
</div>

Finally, let's take a look at the option’s end-of-period (maturity) inner values.

<div style= "text-align:center">
<a href="/img/posts/montecarlo/end_inner_hist.png" target="_blank"><img src="/img/posts/montecarlo/end_inner_hist.png"  alt="Sharpen" style="width:50%; margin:1%"></a>
</div>

As you can see, the majority of the simulated values are zero, indicating that the European call option expires worthless in these cases.

## Conclusion

---

We have seen basic examples and how one can use the Monte Carlo method to find answers to certain problems. We also have seen two major applications, the numerical integration and how to estimate an option price using the Black-Scholes-Merton model.

By now, you should've realized that the Monte Carlo method gives you immense problem-solving powers, even if you're not very familiar with the underlying theory or even if such a theory doesn't exist. For instance, see the <a href="https://rstudio-pubs-static.s3.amazonaws.com/241232_eebe419a0aaa4eb89398ee2a61ad3dc2.html" target="_blank">percolation problem</a>, where no mathematical solution for determining the percolation threshold $p^{\ast}$ has yet been derived.

Now you can successfully apply this technique to your problems and become a practitioner of the art of solving problems using Monte Carlo simulations! Good Luck!

## Recommended Reading

- [Go Programming Language 1st Edition - Donovan & Kernighan](https://amzn.to/32zSNNN)

- [Go in Action 1st Edition  -  William Kennedy, Brian Ketelsen, Erik St. Martin](https://amzn.to/2P8ZFP8)

- [Learning Go: An Idiomatic Approach to Real-World Go Programming 1st Edition - Jon Bodner](https://amzn.to/3aq87Ru)

- [Essentials of Monte Carlo Simulation: Statistical Methods for Building Simulation Models 2013th Edition  - Nick T. Thomopoulos](https://amzn.to/3n9gyWy)

- [Options, Futures, and Other Derivatives 10th Edition - John C Hull](https://amzn.to/3v6vbN2)

- [Financial Calculus: An Introduction to Derivative Pricing 1st Edition - Martin Baxter and Andrew Rennie](https://amzn.to/3dWysst)

- [Monte Carlo Methods in Financial Engineering (Stochastic Modelling and Applied Probability - Paul Glasserman](https://amzn.to/3sDOI5O)

- [Python for Finance: Mastering Data-Driven Finance 2nd Edition - Yves Hilpisch](https://amzn.to/3vfANVv)

<div style= "text-align:center; margin: 2%">
<a href="https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440?dchild=1&keywords=go+programming&qid=1619031216&sr=8-1&linkCode=li2&tag=ggcarvalho-20&linkId=4eda68e7e635de0acfe6d799ac104eea&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=0134190440&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=0134190440" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Go-Action-William-Kennedy/dp/1617291781?dchild=1&keywords=go+in+action&qid=1619031323&sr=8-1&linkCode=li2&tag=ggcarvalho-20&linkId=61bc88bddd4d83cf9005d0997fa4416e&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=1617291781&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=1617291781" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Learning-Go-Idiomatic-Real-World-Programming-ebook/dp/B08XYGCM71?dchild=1&keywords=Learning+Go&qid=1616459752&sr=8-2&linkCode=li2&tag=ggcarvalho-20&linkId=043a1f18d943b96ccfa904665331974a&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=B08XYGCM71&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=B08XYGCM71" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Essentials-Monte-Carlo-Simulation-Statistical/dp/1489986081?dchild=1&keywords=monte+carlo+simulation&qid=1619031483&sr=8-1&linkCode=li2&tag=ggcarvalho-20&linkId=cbeb4e0c855e7978f05b9961222e2c3c&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=1489986081&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=1489986081" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Options-Futures-Other-Derivatives-Tenth/dp/9352866592?dchild=1&keywords=john+c+hull&qid=1619031541&sr=8-1&linkCode=li2&tag=ggcarvalho-20&linkId=cbd54d235a891040add4a1252ef53758&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=9352866592&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=9352866592" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Financial-Calculus-Introduction-Derivative-Pricing/dp/0521552893?dchild=1&keywords=Financial+Calculus%3A+An+Introduction+to+Derivative+Pricing&qid=1619543170&sr=8-1&linkCode=li2&tag=ggcarvalho-20&linkId=3be2d3ba804e3b91057442a664098a33&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=0521552893&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=0521552893" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Financial-Engineering-Stochastic-Modelling-Probability/dp/0387004513?dchild=1&keywords=monte+carlo+simulation&qid=1619031483&sr=8-7&linkCode=li2&tag=ggcarvalho-20&linkId=9aadabc08118c855b9ee8176857c90c5&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=0387004513&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=0387004513" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
<a href="https://www.amazon.com/Python-Finance-Mastering-Data-Driven/dp/1492024333?crid=39PRDWB4NE6UI&dchild=1&keywords=yves+hilpisch&qid=1619037117&sprefix=yves+h%2Caps%2C245&sr=8-4&linkCode=li2&tag=ggcarvalho-20&linkId=d79beb0dfb8a9e6d06acbb038e10642c&language=en_US&ref_=as_li_ss_il" target="_blank"><img border="0" src="//ws-na.amazon-adsystem.com/widgets/q?_encoding=UTF8&ASIN=1492024333&Format=_SL160_&ID=AsinImage&MarketPlace=US&ServiceVersion=20070822&WS=1&tag=ggcarvalho-20&language=en_US" ></a><img src="https://ir-na.amazon-adsystem.com/e/ir?t=ggcarvalho-20&language=en_US&l=li2&o=1&a=1492024333" width="1" height="1" border="0" alt="" style="border:none !important; margin:0px !important;" />
</div>

By clicking and buying any of these from Amazon after visiting the links above, I might get a commission from their <a href="https://affiliate-program.amazon.com/" target="_blank">Affiliate program</a>, and you will be contributing to the growth of this blog :)