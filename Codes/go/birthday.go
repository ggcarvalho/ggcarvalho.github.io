package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	num_people := 23
	trials := 1000000
	sucess := 0

	for i := 0; i < trials; i++ {
		bdays := gen_bday_list(num_people)
		uniques := unique(bdays)

		if !(len(bdays)==len(uniques)) {
			sucess++
		}
	}

	probability := float64(sucess) / float64(trials)
	fmt.Printf("The probability of at least 2 persons in a group of %d people share a birthday is %.2f%%\n", num_people, 100.0*probability)
}

// returns a slice with the unique elements of a given slice
func unique(s []int) []int {
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
func gen_bday_list(n int) []int {
	var bdays []int
	for i := 0; i < n; i++ {
		bday := rand.Intn(365)
		bdays = append(bdays, bday)
	}
	return bdays
}