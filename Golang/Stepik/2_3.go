package main

import "fmt"

func main() {
	fmt.Println("Привет, как вас зовут?")
	fmt.Println("Введите 3 имени")
	var name1, name2, name3 string
	// fmt.Scan(&name1, &name2, &name3)
	_, _ = fmt.Scan(&name1, &name2, &name3)
}
