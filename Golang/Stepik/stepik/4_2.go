package main

import "fmt"

func main() {
	var c, n int
	fmt.Scan(&n)
	sum := 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&c)
		if c == 0 {
			sum++
		}
	}
	if sum != 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
