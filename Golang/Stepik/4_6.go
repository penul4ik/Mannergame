package main

import (
	"fmt"
)

func main() {
	n := 1
	last := 1
	var cnt int
	for n != 0 {
		if last < n {
			cnt++
		}
		last += n
		fmt.Scan(&n)
	}
	fmt.Println(cnt)
}
