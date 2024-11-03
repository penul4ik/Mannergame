package main

import "fmt"

func main() {
	var l int
	fmt.Scan(&l)
	var array = make([]int, l)
	for i := 0; i < len(array); i++ {
		fmt.Scan(&array[i])
		if i%2 != 0 {
			tmp := array[i-1]
			array[i-1] = array[i]
			array[i] = tmp
		}
	}
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
}
