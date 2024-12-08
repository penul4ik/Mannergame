package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	_ = scan.Scan()
	first := scan.Text()
	_ = scan.Scan()
	second := scan.Text()

	if len(first) == 0 || len(second) == 0 {
		fmt.Println("NO")
	} else {

		fmt.Println(first[len(first)-1], second[0])
	}

}
