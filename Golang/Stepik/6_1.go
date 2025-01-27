package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	inputRune := []rune(input)
	slice := make([]rune, len(input))
	for n, inputSym := range inputRune {
		switch {
		case n != 0 && inputSym == 46 && inputRune[n-1] == 45:
			slice[n] = '1'
			inputRune[n], inputRune[n-1] = 0, 0
		case inputSym == 46:
			slice[n] = '0'
		case n != 0 && inputSym == 45 && inputRune[n-1] == 45:
			slice[n] = '2'
			inputRune[n], inputRune[n-1] = 0, 0
		default:
			continue
		}
	}

	fmt.Println(string(slice))
}
