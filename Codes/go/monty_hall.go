package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	trials := 10000000
	fmt.Printf("Estimating the probability of winning by switching doors with %d game(s).\n\n", trials)

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
	guest_door := rand.Intn(3)
	prize_door := rand.Intn(3)
	goat1 := true
	goat2 := true

	var montys_choice int
	var new_door int
	var goat1_door int
	var goat2_door int
	var switch_door bool
	var show_goat bool

	for goat1 {
		goat1_door = rand.Intn(3)
		if goat1_door != prize_door {
			goat1 = false
		}
	}

	for goat2 {
		goat2_door = rand.Intn(3)
		if (goat2_door != prize_door) && (goat2_door != goat1_door) {
			goat2 = false
		}
	}

	switch_door = true
	show_goat = true

	for show_goat {
		montys_choice = rand.Intn(3)
		if (montys_choice != prize_door) && (montys_choice != guest_door) {
			show_goat = false
		}
	}

	for switch_door {
		new_door = rand.Intn(3)
		if (new_door != guest_door) && (new_door != montys_choice) {
			switch_door = false
		}
	}
	return new_door, prize_door
}