package main

import (
	"fmt"
)

func main() {
	var number, cnt int
	_, _ = fmt.Scan(&number)
	for number != 0 {
		if number%2 == 1 {
			cnt++
		}
		number /= 2
	}
	fmt.Println(cnt)
}
