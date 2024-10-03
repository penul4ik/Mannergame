package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	name := scanner.Text()
	scanner2 := bufio.NewScanner(os.Stdin)
	_ = scanner2.Scan()
	name2 := scanner2.Text()
	scanner3 := bufio.NewScanner(os.Stdin)
	_ = scanner3.Scan()
	name3 := scanner3.Text()
	fmt.Print(name, "\n", name2, "\n", name3)
}
