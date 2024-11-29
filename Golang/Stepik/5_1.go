package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != i && slice[i]+slice[j] == 0 {
				fmt.Print(i, " ")
			}
		}
	}
}
