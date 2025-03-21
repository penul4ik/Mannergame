package main

import "fmt"

func main() {
	var month int
	fmt.Scan(&month)
	season := "Зима"
	switch month {
	case 3, 4, 5:
		season = "Весна"
	case 9, 10, 11:
		season = "Осень"
	case 6, 7, 8:
		season = "Лето"
	}
	fmt.Println(season)
}
