// main.go
package main

import (
	"fmt"
)

type Handler struct {
	divisor int
	text    string
}

var ClassicHandlers = []Handler{
	{3, "fizz"},
	{5, "buzz"}}

func handleNumber(i int) (s string) {
	for _, h := range ClassicHandlers {
		if i%h.divisor == 0 {
			s += h.text
		}
	}
	if len(s) == 0 {
		s = fmt.Sprint(i)
	}
	return
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(handleNumber(i))
	}
}
