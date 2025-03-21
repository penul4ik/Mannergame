package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	cnt := 0
	for x := 0; x <= 10; x++ {
		if x-e == 0 {
			continue
		} else if (a*(x*x*x)+b*(x*x)+c*x+d)/(x-e) == 0 && (a*(x*x*x)+b*(x*x)+c*x+d)%(x-e) == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
