package main

import "fmt"

func main() {
	message := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	runes := []rune(message)
	for i, keywordV := range keyword {
		fmt.Printf("%v %c %v %c ", keywordV, keywordV, runes[i], runes[i])
		if runes[i]-keywordV < 65 {
			b := runes[i] - keywordV + 26
			fmt.Printf("%v %c\n", b, b)
		} else {
			fmt.Print("да")
		}
	}
	fmt.Printf("%v %c %v %c\n", 'A', 65, 'Z', 90)
}
