package main

import (
	"fmt"
)

func main() {
	var two, three, five, six int
	fmt.Scan(&two, &three, &five, &six)
	var max256 *int
	if two <= five && two <= six {
		max256 = &two
	} else if five <= two && five <= six {
		max256 = &five
	} else {
		max256 = &six
	}
	rMax256 := *max256
	two -= *max256
	var max32 *int
	if two <= three {
		max32 = &two
	} else {
		max32 = &three
	}
	fmt.Println((256 * rMax256) + (32 * *max32))
}
