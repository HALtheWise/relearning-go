// handlers.go
package main

import (
	"fmt"
)

type DivisorHandler struct {
	divisor int
	text    string
}

func (d DivisorHandler) process(i int, input string) string {
	if i%d.divisor == 0 {
		return input + d.text
	} else {
		return input
	}
}

type NumberHandler struct{}

func (d NumberHandler) process(i int, input string) string {
	if len(input) == 0 {
		return fmt.Sprint(i)
	} else {
		return input
	}
}
