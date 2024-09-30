package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	Name string
}

func (h *Human) WhoAmI() {
	fmt.Printf("I am %s\n", h.Name)
}

func (h *Human) Feed(food string, times int) {
	fmt.Printf("I am %s and i eat the %s %d times\n", h.Name, food, times)
}

func IsHuman(h *Human) bool {
	return true
}

func main() {
	// fmt.Println("Hello, World!")
	http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))

	human := &Human{
		Name: "Makson",
	}

	human.WhoAmI()

	human.Feed("apple", 10)

	if IsHuman(human) {
		fmt.Println("bazinga")
	}
}
