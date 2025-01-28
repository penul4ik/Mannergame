package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//startDay := 30
	//startMouth := 10
	//startYear := 2020
	//distance := 62100000
	fmt.Printf("Spaceline           Days  Trip type  Price\n")
	fmt.Printf("==========================================\n")
	for i := 1; i <= 10; i++ {
		CompanyName := "None"
		switch {
		case rand.Intn(4)+1 == 1:
			CompanyName = "VirginGalactic"
		default:
			CompanyName = "SpaceX"
		}
		fmt.Printf("%-12v %11v %8v    $ %3v\n", CompanyName, "10", "One-way", "100")
	}
}
