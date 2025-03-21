package main

import "fmt"

func main() {
	a := 9000
	b := 9050
	for i := a; i <= b; i++ {
		d4 := i % 10
		d3 := i % 100 / 10
		d2 := i % 1000 / 100
		d1 := i % 10000 / 1000
		//fmt.Println(d1, d2, d3, d4)
		if d1 != d2 && d2 != d3 && d3 != d4 && d1 != d4 && d2 != d4 && d1 != d3 && i != a {
			fmt.Println(i)
			break
		}
	}
}
