package main

import (
	"fmt"
	"time"
	"math/rand"
)

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

func gen_bday_list(n int) []int {
	var bdays []int
	for i := 0; i < n; i++ {
		bday := rand.Intn(365) + 1
		bdays = append(bdays, bday)
	}
	return bdays
}

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