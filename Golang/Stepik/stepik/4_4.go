package main

import (
	"fmt"
)

func main() {
	var a, b, x, y int
	fmt.Scan(&a, &b)
	remOfDiv := 1
	if a > b {
		x = a
		y = b
	} else {
		x = b
		y = a
	}
	if y == 0 {
		fmt.Println(x)
	} else {
		for remOfDiv != 0 {
			remOfDiv = x % y
			x = y
			y = remOfDiv
			fmt.Println(x, y, remOfDiv)
		}
		fmt.Println(x)
	}
}
