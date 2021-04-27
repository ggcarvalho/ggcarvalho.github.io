package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	numGames := 10_000_000
	fmt.Printf("Estimating the probability of winning by switching doors with %d game(s).\n\n", numGames)

	sucess := 0
	for i := 0; i < numGames; i++ {
		newDoor, prizeDoor := set_monty_hall()
		if newDoor == prizeDoor {
			sucess++
		}
	}
	probability := float64(sucess) / float64(numGames)
	theoreticalValue := 2.0 / 3.0

	errorPct := 100.0*math.Abs(probability - theoreticalValue) / theoreticalValue

	fmt.Printf("Estimated probability: %9f \n", probability)
	fmt.Printf("Theoretical value: %9f \n", theoreticalValue)
	fmt.Printf("Error: %9f%%\n", errorPct)
}

// randomly sets the game
func set_monty_hall() (int, int) {
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
		if  (prizeDoor != goat1Door && prizeDoor != goat2Door && goat1Door != goat2Door) {areDoorsSelected = true}
	}

	showGoat := false
	for !showGoat {
		montysChoice = rand.Intn(3)
		if (montysChoice != prizeDoor) && (montysChoice != guestDoor) {
			showGoat = true
		}
	}

	madeSwitch := false
	for !madeSwitch {
		newDoor = rand.Intn(3)
		if (newDoor != guestDoor) && (newDoor != montysChoice) {
			madeSwitch = true
		}
	}
	return newDoor, prizeDoor
}