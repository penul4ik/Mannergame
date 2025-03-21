package main

import (
	"fmt"
)

func main() {
	const minInHour int = 60
	var min int
	fmt.Scan(&min)
	result_hour := min / minInHour
	result_min := min % minInHour
	fmt.Println(min, "мин - это", result_hour, "час", result_min, "минут.")
}
