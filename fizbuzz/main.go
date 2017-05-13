// main.go
package main

import (
	"fmt"
)

type Handler interface {
	process(i int, input string) (output string)
}

var ClassicHandlers = []Handler{
	DivisorHandler{3, "fizz"},
	DivisorHandler{5, "buzz"},
	NumberHandler{},
}

func handleNumber(n int, handlers []Handler) (s string) {
	for _, h := range handlers {
		s = h.process(n, s)
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
