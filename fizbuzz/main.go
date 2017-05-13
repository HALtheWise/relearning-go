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

func handleNumber(i int, handlers []Handler) (s string) {
	for _, h := range handlers {
		if i%h.divisor == 0 {
			s += h.text
		}
	}
	if len(s) == 0 {
		s = fmt.Sprint(i)
	}
	return
}

func RunFizzBuzz(start, end int, handlers []Handler) {
	for i := start; i <= end; i++ {
		fmt.Println(handleNumber(i, handlers))
	}
}

func main() {
	RunFizzBuzz(1, 100, ClassicHandlers)
}
