package main

import (
	"fmt"
)

func sign(n int) bool {
	a := n / 1000
	b := n % 1000
	z := a/100 + (a%100)/10 + a%10
	y := b/100 + (b%100)/10 + b%10
	if y == z {
		return true
	} else {
		return false
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if sign(a) && sign(b) {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
