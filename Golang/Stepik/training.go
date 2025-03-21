package main

import "fmt"

func increment(n *int) {
	// *n обращается к значению, на которое указывает n, и увеличивает его на 1
	*n = *n + 1
}

func main() {
	a := 10
	fmt.Println("До:", a) // До: 10

	// Передаём адрес переменной a в функцию
	increment(&a)

	fmt.Println("После:", a) // После: 11
}
