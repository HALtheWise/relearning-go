// grid
package main

import (
	"fmt"
	"strings"
)

type Grid struct {
	FixedValues    [9][9]int8
	PossibleValues [9][9][9]bool
}

func (g *Grid) String() string {
	var lines []string
	for row := 0; row < 9; row++ {
		var s []string
		for col := 0; col < 9; col++ {
			val := g.FixedValues[row][col]
			if val == 0 {
				s = append(s, "*")
			} else {
				s = append(s, fmt.Sprint(val))
			}
		}
		lines = append(lines, strings.Join(s, ""))
	}
	return strings.Join(lines, "\n")
}
