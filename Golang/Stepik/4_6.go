package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	last := n
	var cnt int
	for n != 0 {
		if last < n {
			cnt++
		}
		last = n
		fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
