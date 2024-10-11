package main

import (
	"fmt"
	"math"
)

func main() {
	var g1, v1, g2, v2 float64
	var a bool = true
	fmt.Scan(&g1, &v1, &g2, &v2)
	if g1 == g1 && v1 == v2 {
		a = false
	}
	if a == true && (g1-v1 == g2-v2) || a == true && (math.Abs(g1-g2) == math.Abs(v1-v2)) {
		fmt.Println("YES")
		fmt.Println(a)
		fmt.Println(g1-v1, g2-v2)
		fmt.Println(g1-g2, v1-v2)
	} else {
		fmt.Println("NO")
		fmt.Println(a)
		fmt.Println(g1-v1, g2-v2)
		fmt.Println(g1-g2, v1-v2)
	}
}
