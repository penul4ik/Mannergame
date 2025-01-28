package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int
	fmt.Println("Угадываем число от 0 до 20")
	fmt.Scan(&num)
	count := 20
	for {
		if rand.Intn(count) == num {
			fmt.Println("Загаданное числo:", count, num)
			break
		}
	}
}
