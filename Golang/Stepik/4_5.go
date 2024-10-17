package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	cnt := 0
	i := a
	for i%3 == 0 {
		i = a / 3
		cnt++
		fmt.Println(cnt, i, a)
	}
	fmt.Println(cnt, 1%3)
}
