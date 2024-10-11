package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := a / 1000
	c := (a / 100) % 10
	d := (a % 100) / 10
	e := a % 10
	fmt.Println(a, b, c, d, e)
	if b == e && c == d {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
