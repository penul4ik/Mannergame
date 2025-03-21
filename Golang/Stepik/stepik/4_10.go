package main

import "fmt"

func main() {
	var i, max, last int = 1, 1, 1
	for i != 0 {
		fmt.Scan(&i)
		if i > max || i > last {
			if i < max {
				last = i
				continue
			} else if i > max {
				last = max
			}
			max = i
		}
	}
	fmt.Println(max, last)
}
