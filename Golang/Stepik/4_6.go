package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	l := n
	var cnt int
	for n != 0 {
		if (l < 0 && n > 0) || (l > 0 && n < 0) {
			cnt++
		}
		l = n
		fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
