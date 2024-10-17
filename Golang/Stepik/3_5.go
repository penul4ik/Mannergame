package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	way1 := a + b + c
	way2 := a*2 + b*2
	way3 := a*2 + c*2
	way4 := b*2 + c*2
	if way1 <= way2 {
		if way1 <= way3 {
			if way1 <= way4 {
				fmt.Println(way1)
			} else {
				fmt.Println(way4)
			}
		} else if way3 <= way4 {
			fmt.Println(way3)
		} else {
			fmt.Println(way4)
		}
	} else if way2 <= way3 {
		if way2 <= way4 {
			fmt.Println(way2)
		} else {
			fmt.Println(way4)
		}
	} else if way3 <= way4 {
		fmt.Println(way3)
	} else {
		fmt.Println(way4)
	}
}
