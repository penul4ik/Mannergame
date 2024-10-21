package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := 1
	k := 0
	for sum <= n {
		if sum*2 > n {
			goto Exit
		}
		sum <<= 1
		k++
	}
Exit:
	fmt.Println(k)
}
