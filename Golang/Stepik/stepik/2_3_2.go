package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	d := *b
	*a = d
	*b = c
}

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
