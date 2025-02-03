package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg"
	for i, _ := range message {
		if message[i] >= 'A' && message[i] <= 'z' {
			fmt.Printf("%c", message[i]-3)
		} else {
			fmt.Printf("%c", message[i])
		}
	}
	fmt.Print("\n")
	spanishMessage := "Hola Estación Espacial Internacional"
	for _, c := range spanishMessage { // Используем rune, а не байты
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c) // Теперь корректно обрабатывает Unicode
	}
}
