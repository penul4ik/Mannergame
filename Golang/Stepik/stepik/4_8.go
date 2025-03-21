package main

import "fmt"

func main() {
	var n, i string
	_, _ = fmt.Scan(&n)
	for cnt := 0; ; {
		cnt++
		fmt.Scan(&i)
		if i == n {
			fmt.Println(cnt)
			break
		}
	}
}
