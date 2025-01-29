package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const goal = 20.00
	const five = 0.05
	const ten = 0.10
	const twentyFive = 0.25
	var sum float64
	for sum < goal {
		switch {
		case rand.Intn(3) == 2:
			fmt.Printf("В копилке %.2f, прибавляем %v\n", sum, five)
			sum += five
		case rand.Intn(3) == 1:
			fmt.Printf("В копилке %.2f, прибавляем %v\n", sum, ten)
			sum += ten
		default:
			fmt.Printf("В копилке %.2f, прибавляем %v\n", sum, twentyFive)
			sum += twentyFive
		}
	}
	fmt.Printf("Накопили %.2f:", sum)
}
