// main.go
package main

import (
	"fmt"
)

func handleNumber(i int) (s string) {
	if i%3 == 0 {
		s += "fizz"
	}
	if i%5 == 0 {
		s += "buzz"
	}
	if len(s) == 0 {
		s += fmt.Sprint(i)
	}
	return
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(handleNumber(i))
	}
}
