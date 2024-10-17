package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		fmt.Println(i * i)
	}
}
