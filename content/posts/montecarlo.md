---
title: "The joy of solving problems with Monte Carlo simulations"
type: post
description: "Using the power of randomness to answer scientific questions."
date: "2021-05-19"
image: "/img/posts/montecarlo/roulette.jpg"
tag: "go"
name: "go"
hashtags: "#go #golang #scientificcomputing #montecarlo #trade #algotrade #quant #finance"
draft: false
---

In this article, we will explore some examples and applications of Monte Carlo simulations using the Go programming language. To keep this article interactive, after each Go code provided you will find a link to the <em>Go Playground</em>, where you can run it without installing Go on your machine.

Put your adventure helmets on!

<div style="text-align:left"><img src="/img/posts/montecarlo/gopher_adventurer.jpg" style="width: 20%"></div>

{{< table_of_contents >}}

## What is Monte Carlo simulation?

---

Generally speaking, Monte Carlo methods consist of a broad class of computational algorithms that rely on repeated random sampling to obtain numerical results. This technique is used throughout in areas such as physics, finance, engineering, project management, insurance, and transportation, where a numerical result is needed and the underlying theory is difficult and/or unavailable. It was invented by [John von Neumann](https://en.wikipedia.org/wiki/John_von_Neumann) and [Stanisław  Ulam](https://en.wikipedia.org/wiki/Stanislaw_Ulam) during World War II to improve decision making under uncertain conditions. It was named after a well-known casino town, called Monaco, since the element of chance is core to the modeling approach, similar to a game of roulette.

If you are new to the subject, you may find it useful to read this [IBM article](https://www.ibm.com/cloud/learn/monte-carlo-simulation) and get acquainted with the concept first.

>"Unlike a normal forecasting model, Monte Carlo Simulation predicts a set of outcomes based on an estimated range of values versus a set of fixed input values. In other words, a Monte Carlo Simulation builds a model of possible results by leveraging a probability distribution, such as a uniform or normal distribution, for any variable that has inherent uncertainty. It, then, recalculates the results over and over, each time using a different set of random numbers between the minimum and maximum values. In a typical Monte Carlo experiment, this exercise can be repeated thousands of times to produce a large number of likely outcomes." [IBM](https://www.ibm.com/cloud/learn/monte-carlo-simulation)

## First examples

---

Let's start our journey with some basic, and even textbook, examples of Monte Carlo simulations. You're going to notice a pattern when implementing this method. Use it to create or try your examples.

### Estimating $\pi$

This is by far the most famous example of Monte Carlo simulations, considered to be the "zeroth example" of the subject. To estimate $\pi$ we need to pose the problem in probabilistic terms. We do so by considering a circle of radius $r=1$ inscribed in a square of side $l=2$, both centered in the origin of a cartesian coordinate system. This situation is depicted below.

<div style="text-align:center"><img src="/img/posts/montecarlo/circle.png" style="width: 40%"></div>

If we were to draw random points in this square, some will fall within the circle and some won't. But the ratio between points inside the circular region and the total amount of points we draw will be closer and closer to the ratio between the area of the circle and the area of the square as we draw more of these random points.

<div style="text-align:center"><img src="/img/posts/montecarlo/circle2.png" style="width: 40%"></div>

As you might know, the area of the circular region is $A_{\circ} = \pi\cdot r^2 = \pi$, since $r = 1$, and the area of the square is $A_{\square} = l^2 = 4$, since $l=2$. Thus, $$\pi = 4\cdot \frac{A_{\circ}}{A_{\square}}.$$

As a result, we can estimate $\pi$ as
$$\pi \approx 4\cdot \frac{\text{$\\#$ points inside the circle}}{\text{$\\#$ points}}.$$
Remember that, in our case, $(x, y)$ will fall inside the circular region if $x^2 + y^2 < 1$.

Now, let's turn this idea into a Go code.

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

	trials := 1000000
	fmt.Printf("Estimating pi with %d point(s).\n\n", trials)

	sucess := 0
	for i := 0; i < trials; i++ {
		px := 2.0*rand.Float64() - 1.0
		py := 2.0*rand.Float64() - 1.0

		if inside_circle(px, py) {
			sucess += 1
		}
	}

	pi_approx := 4.0*(float64(sucess) / float64(trials))
	pi := math.Pi
	fmt.Printf("Estimated pi: %9f \n", pi_approx)
	fmt.Printf("pi: %9f \n", pi)

	error_pct := 100*abs(pi_approx - pi) / pi

	fmt.Printf("Error: %9f%%\n", error_pct)
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
```

[Run this code in The Go Playground](https://play.golang.org/)

```bash
$ go run pi.go

Estimating pi with 1000000 point(s).

Estimated pi:  3.143056
pi:  3.141593
Error:  0.046580%
```

We were able to approximate $\pi$ with an error of $0.053965$% ! Let's jump to the next exmaple.

### Estimating Euler's number

Not long ago, Lex Fridman published the following in a LinkedIn post:

<div style="text-align:center"><img src="/img/posts/montecarlo/lex.png" style="width: 60%"></div>

This is very intriguing! Since I have no idea how to prove this statement, I decided to write a simple Python program to test it myself. I wrote my code as a response to Lex's post, and people were surprised by it (this is the actual motivation to write this blog post).
This, once again, showed me the power of Monte Carlo simulations in scientific computing, where you can find precise numerical answers to your problems without relying on any theoretical background (they might not even exist!).

I wrote an equivalent Go code to solve this problem. Here it is:

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

	trials := 10000000
	fmt.Printf("Estimating e with %d trial(s).\n\n", trials)

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

// absolute value of x
func abs(x float64) float64 {
	if x < 0.0 {
		return -x
	}
	return x
}
```

[Run this code in The Go Playground](https://play.golang.org/)

```bash
$ go run euler.go

Estimating e with 10000000 trial(s).

Expected vale:  2.718518
e:  2.718282
Error:  0.008681%
```

An astonishing result!

### The birthday paradox

This is a famous problem in statistics:
>In a group of $23$ people, the probability of a shared birthday exceeds $50$%.

That sounds weird at first, but if you're good enough with math you can easily prove this statement. However, we are not interested in formal proofs here. That is the whole point of these simulations. The idea is very simple: create a list with $n$ ($23$, in our case) random numbers between $1$ and $365$ representing the birth day of each person and if (at least) two of them coincide, you increment the `success` variable. Do it a certain number of times and calculate the probability dividing the number of sucesses by the total number of simulations.

The corresponding Go code:

```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	trials := 1000000
	sucess := 0

	for i := 0; i < trials; i++ {

		bdays := gen_bday_list(23)
		uniques := unique(bdays)

		if !(len(bdays)==len(uniques)) {
			sucess++
		}
	}

	probability := float64(sucess) / float64(trials)

	fmt.Printf("The probability of at least 2 persons in a group of 23 people share a birthday is %.2f%%\n", 100*probability)

}

// returns a slice with the unique elements of a given slice
func unique(intSlice []int) []int {
    keys := make(map[int]bool)
    list := []int{}
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }
    return list
}

// generates the list of birth days
func gen_bday_list(n int) []int {
	var bdays []int
	for i := 0; i < n; i++ {
		bday := rand.Intn(365) + 1
		bdays = append(bdays, bday)
	}
	return bdays
}
```

[Run this code in The Go Playground](https://play.golang.org/)

```bash
$ go run birthday.go

The probability of at least 2 persons in a group of 23 people share a birthday is 50.73%
```

You can tweak this code to reproduce the following table from [Wikipedia](https://en.wikipedia.org/wiki/Birthday_problem):

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

Of course, for $n \geq 365$ you don't need any calculations, it's a straightforward consequence of the [<em>pigeonhole principle</em>](https://en.wikipedia.org/wiki/Pigeonhole_principle).

## The (in)famous Monty Hall problem

---

That is a problem that has been disturbing people for ages. Just like the birthday problem, you can solve it using basic math/probability theory, which we won't do. Let's state the problem and provide a Monte Carlo simulation to solve it.

The problem is:
>Suppose you're on a game show, and you're given the choice of three doors: Behind one door is a valuable prize; behind the others, goats. You pick a door, say No. $1$, and the host, who knows what's behind the doors, opens another door, say No. $3$, which has a goat. He then says to you, "Do you want to pick door No. $2$?" Is it to your advantage to switch your choice?

<div style="text-align:center"><img src="/img/posts/montecarlo/montyhall.png" style="width: 60%; margin: 2%"></div>

People often believe that they are in a $50-50$ situation and therefore the switch is not very relevant. But if you carefully solve this problem, you will find that there is a probability of $2/3 \approx 66.7$% to win the prize if you decide to switch doors.

The following Go code simulates this game and estimates the probability of winning if the guest chooses to switch doors. Note that we first set the game so that the doors are properly chosen (possibly not an optimal code), and then we simulate the game to estimate the desired probability.

```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	trials := 10000000
	fmt.Printf("Estimating the propability of winning by switching doors with %d trial(s).\n\n", trials)

	sucess := 0
	for i := 0; i < trials; i++ {
		new_door, prize_door := set_monty_hall()
		if new_door == prize_door {
			sucess++
		}
	}
	probability := float64(sucess) / float64(trials)
	theoretical_value := 2.0 / 3.0

	error_pct := 100*abs(probability - theoretical_value) / theoretical_value

	fmt.Printf("Estimated probability: %9f \n", probability)
	fmt.Printf("Theoretical value: %9f \n", theoretical_value)
	fmt.Printf("Error: %9f%%\n", error_pct)
}

// absolute value of x
func abs(x float64) float64 {
	if x < 0.0 {
		return -x
	}
	return x
}

// randomly sets the game
func set_monty_hall() (int, int) {
	guest_door := rand.Intn(3) + 1
	prize_door := rand.Intn(3) + 1
	goat1 := true
	goat2 := true

	var montys_choice int
	var new_door int
	var goat1_door int
	var goat2_door int
	var switch_door bool
	var show_goat bool

	for goat1 {
		goat1_door = rand.Intn(3) + 1
		if goat1_door != prize_door {
			goat1 = false
		}
	}

	for goat2 {
		goat2_door = rand.Intn(3) + 1
		if (goat2_door != prize_door) && (goat2_door != goat1_door) {
			goat2 = false
		}
	}

	switch_door = true
	show_goat = true

	for show_goat {
		montys_choice = rand.Intn(3) + 1
		if (montys_choice != prize_door) && (montys_choice != guest_door) {
			show_goat = false
		}
	}

	for switch_door {
		new_door = rand.Intn(3) + 1
		if (new_door != guest_door) && (new_door != montys_choice) {
			switch_door = false
		}
	}
	return new_door, prize_door
}
```

[Run this code in The Go Playground](https://play.golang.org/)

```bash
$ go run monty_hall.go

Estimating the probability of winning by switching doors with 10000000 trial(s).

Estimated probability:  0.666596
Theoretical value:  0.666667
Error:  0.010540%
```

Therefore, contrary to popular belief, it is more advantageous to the guest to switch doors confirming the theoretical result.

## Integration using Monte Carlo simulations

---

Now, let's see how we can use the Monte Carlo method to find the value of definite integrals of continuous functions in a specified range.

Just as a reminder, if $f: [a,b] \rightarrow \mathbb{R}$ is a continuous function, then the quantity $$S = \int_a^b f(x)dx,$$
represents the area of the region between the graph of $f$ and the $x-$axis.

<div style="text-align:center"><img src="/img/posts/montecarlo/integralarea.png" style="width: 30%; margin: 2%"></div>

One important feature of this property is that this area has a sign, having negative values if the region is below the $x-$axis.

One way to aprroximate this area is  $$S \approx \frac{b-a}{n}\sum_{i=1}^n f(a + (b-a)U_i),$$
where $U_i \sim \mathcal{U}(0,1)$. Feel free to try different probability distributions and compare the results.

We are going to use this technique to solve a classic problem. If you are a calculus geek, you might know how difficult it is to calculate the integral $$S = \int_{-\infty}^{\infty} e^{-x^2}dx.$$

It involves a trick using Fubini's theorem and change in from cartesian to polar coordinates. Surprisingly, the result of this integral is $\sqrt{\pi}$. Let's use Monte Carlo integration to evaluate $$\bar{S} = \int_{-20}^{20} e^{-x^2}dx.$$

<div style="text-align:center"><img src="/img/posts/montecarlo/function.png" style="width: 35%; margin: 2%"></div>

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

	trials := 1000000
	fmt.Printf("Estimating the integral of f with %d point(s).\n\n", trials)

	integral := monte_carlo_integral(gaussian, -20.0, 20.0, trials)
	fmt.Printf("Approx. integral: %9f \n", integral)
}

// MC integrator
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

// function to be integrated
func gaussian(x float64) float64 {
	return math.Exp(-x*x)
}
```

[Run this code in The Go Playground](https://play.golang.org/)

```bash
$ go run mc_integration.go

Estimating the integral of f with 1000000 point(s).

Approx. integral:  1.772819
```

In fact, $1.772819^2 \approx 3.143$.

## Option pricing using the Black-Scholes model

---

Ok, for the final section of this article I have something special that draws the attention of many people around the world.

The Black–Scholes, or Black–Scholes–Merton model, is a mathematical model for the dynamics of a financial market containing derivative investment instruments, giving a theoretical estimate of the price of <em>European-style options</em> and shows that the option has a unique price given the risk of the security and its expected return. This work granted Fischer Black and Myron Scholes their Nobel Prize in economics and has been widely used in algorithmic trading strategies around the world.

Before we move on, I have a <strong>disclaimer</strong>: we will dive into some math!

### The equation

We start with the Black-Scholes-Merton option pricing formula:
\begin{eqnarray*}
C(S_t, K, t, T, r, \sigma) &=& S_t N(d_1) - e^{-r(T-t)}KN(d_2)\newline
N(d) &=& \frac{1}{\sqrt{2\pi}}\int_{-\infty}^d e^{-\frac{1}{2}x^2}dx \newline
d_1 &=& \frac{\log\frac{S_t}{K} + (T-t)\left(r + \frac{\sigma^2}{2}\right)}{\sigma\sqrt{T-t}}\newline
d_2 &=& \frac{\log\frac{S_t}{K} + (T-t)\left(r - \frac{\sigma^2}{2}\right)}{\sigma\sqrt{T-t}}.
\end{eqnarray*}

In the equations above $S_t$ is the price of the underlying asset at time $t$, $\sigma$ is the constant volatility (standard deviation of returns) of the underlying asset, $K$ is the strike price of the option, $T$ is the maturity date of the option, $r$ is the risk-free short rate.

The Black-Scholes-Merton stochastic differential equation is given by $$dS_t = rS_t dt + \sigma S_t dZ_t,$$
where $Z(t)$ is the random component of the model (a brownian motion). We will look at the discretized version of the BSM model (Euler discretization), given by $$S_t = S_{t-\Delta t}  +  \exp\left(\left(r - \frac{\sigma^2}{2}\right)\Delta t + \sigma\sqrt{\Delta t}Z_t \right).$$

We will optimize things by taking the logarithm of the discretized equetion, yielding: $$\log S_t = \log S_{t-\Delta t}  +  \left(r - \frac{\sigma^2}{2}\right)\Delta t + \sigma\sqrt{\Delta t}z_t .$$

The variable z is a standard normally distributed random variable, $0 < \Delta t < T$, a (small
enough) time interval. It also holds $0 < t \leq T$ with $T$ the final time horizon.

In this simulation we use the values $S_0 = 100$, $K = 105$, $T = 1.0$, $r = 0.05$, $\sigma = 0.2$.

#### The Simulation

We follow the steps:

1. Divide the time interval $[0, T]$ in equidistant subintervals of length $\Delta t$.
2. Start iterating $i = 1, 2,..., I$.
    - For every time step $t  \in \\{\Delta t, 2\Delta t,..., T \\}$, draw pseudorandom numbers $z_t(i)$.
    - Determine the time $T$ value of the index level $S_T(i)$ by applying the pseudo-
    random numbers time step by time step to the discretized equation.
    - Determine the inner value $h_T$ of the European call option at $T$ as $h_T(S_T(i)) =
    \max(S_T(i) – K, 0)$.
    - Iterate until $i = I$.

3. Sum up the inner values, average, and discount them back with the riskless short
rate according to the formula
$$C_0 \approx e^{-rT} \frac{1}{I} \sum_I h_T(S_T(i)),$$
called the Monte Carlo estimator for the European call option.

Without any further ado, here is the corresponding Go code:

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
```

[Run this code in The Go Playground](https://play.golang.org/)

```bash
$ go run black_scholes.go

European Option Value: 8.027
Execution time:  430.464289ms
```

Let's compare the results with the same simulation in Python (taken from Yves Hilpisch's "Python for Finance"):

<div style="text-align:center"><img src="/img/posts/montecarlo/yves.png" style="width: 50%; margin: 2%"></div>

Nearly the same result in a fraction of the time! To be completely fair, when the author uses full Numpy vectorization the results are much better in terms of performance, although we still have a clear winner.

<div style="text-align:center"><img src="/img/posts/montecarlo/yves2.png" style="width: 50%; margin: 2%"></div>

For the sake of completeness, let's plot some stuff and have a graphical look at the underlying mechanics.

First, let's plot the simulated index levels (the paths taken during the simulation). The figures below represent the first $10$, the first $100$, and all simulated index levels respectively.

<div style= "text-align:center">
<a href="/img/posts/montecarlo/10.png" target="_blank"><img src="/img/posts/montecarlo/10.png"  alt="Mean" style="width:40%; margin:1%"></a>
<a href="/img/posts/montecarlo/100.png" target="_blank"><img src="/img/posts/montecarlo/100.png"  alt="Gaussian" style="width:40%; margin:1%"></a>
</div>
<div style= "text-align:center">
<a href="/img/posts/montecarlo/all.png" target="_blank"><img src="/img/posts/montecarlo/all.png"  alt="Sharpen" style="width:40%; margin:1%"></a>
</div>

Second, we want to see the frequency of the simulated index levels at the end of the
simulation period. We expect, by inspectioning every path, that it will be normally distributed around a value closer to the initial index level.

<div style= "text-align:center">
<a href="/img/posts/montecarlo/end_hist.png" target="_blank"><img src="/img/posts/montecarlo/end_hist.png"  alt="Sharpen" style="width:50%; margin:1%"></a>
</div>

Finally, let's take a look at the option’s end-of-period (maturity) inner values.

<div style= "text-align:center">
<a href="/img/posts/montecarlo/end_inner_hist.png" target="_blank"><img src="/img/posts/montecarlo/end_inner_hist.png"  alt="Sharpen" style="width:50%; margin:1%"></a>
</div>

As you can see, the majority of the simluated values are zero, indicating that the European call option expires worthless in many cases.

## Conclusion

---

We have seen how one can use the Monte Carlo method find answers to certain problems. We alse have seen two major applications, the numerical integration and how to estimate an option price using the Black-Scholes-Merton model.

## Recommended reading

- [Go Programming Language 1st Edition - Donovan & Kernighan](https://amzn.to/32zSNNN)

- [Go in Action 1st Edition  -  William Kennedy, Brian Ketelsen, Erik St. Martin](https://amzn.to/2P8ZFP8)

- [Learning Go: An Idiomatic Approach to Real-World Go Programming 1st Edition - Jon Bodner](https://amzn.to/3aq87Ru)

- [Essentials of Monte Carlo Simulation: Statistical Methods for Building Simulation Models 2013th Edition  - Nick T. Thomopoulos](https://amzn.to/3n9gyWy)

- [Options, Futures, and Other Derivatives 10th Edition - John C Hull](https://amzn.to/3v6vbN2)

- [Monte Carlo Methods in Financial Engineering (Stochastic Modelling and Applied Probability - Paul Glasserman](https://amzn.to/3sDOI5O)

- [Python for Finance: Mastering Data-Driven Finance 2nd Edition - Yves Hilpisch](https://amzn.to/3vfANVv)

By clicking and buying any of these from Amazon after visiting the links above, I might get a commission from their [Affiliate program](https://affiliate-program.amazon.com/), and you will be contributing to the growth of this blog :)