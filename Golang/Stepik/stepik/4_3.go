package main

import "fmt"

func main() {
	var n int
	m := 1
	_, _ = fmt.Scan(&n)
	for i := 2; i <= n; i += 2 {
		m *= i
	}
	fmt.Println(31 % 3)
}
