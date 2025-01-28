package main

import (
	"fmt"
	"math/rand"
)

func main() {
	distance := 62100000
	minKmPerDays := 1382400
	maxKmPerDays := 2592000
	maxDays := distance / minKmPerDays // 44
	minDays := distance / maxKmPerDays // 23
	fmt.Printf("Spaceline           Days  Trip type     Price\n")
	fmt.Printf("=============================================\n")
	for i := 1; i <= 10; i++ {
		companyName := "None"
		tripType := "None"
		days := 0
		switch {
		case rand.Intn(2)+1 == 1:
			companyName = "VirginGalactic"
		case rand.Intn(2)+1 == 2:
			companyName = "Space Adventures"
		default:
			companyName = "SpaceX"
		}
		switch {
		case rand.Intn(2) == 1:
			tripType = "Round-trip"
		default:
			tripType = "One-way"
		}
		for {
			days = rand.Intn(maxDays) + 1
			if days > minDays {
				break
			}
		}
		fmt.Printf("%-21v %2v  %-10v    $ %3v\n", companyName, days, tripType, days+20)
	}
}
