// grid
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Grid struct {
	FixedValues    [9][9]int8
	PossibleValues [9][9][9]bool
}

func (g *Grid) getConflictValues(i, j int) (conflicts []int8) {
	// TODO fix naming conventions for variables
	for ii := 0; ii < 9; ii++ {
		// Add the values in my column but not in my row
		if i != ii {
			conflicts = append(conflicts, g.FixedValues[ii][j])
		}
	}
	for jj := 0; jj < 9; jj++ {
		// Add the values in my row but not in my column
		if j != jj {
			conflicts = append(conflicts, g.FixedValues[i][jj])
		}
	}

	// Add the values in my box that aren't me
	// TODO this
	return
}

func (g *Grid) Update() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			g.updatePossibilities(i, j)
		}
	}
}

func (g *Grid) updatePossibilities(i, j int) {
	// TODO Be able to do this operation incrementally whenever an element is added to the grid
	// Reset all possiblities to "true"
	for k := 0; k < 9; k++ {
		g.PossibleValues[i][j][k] = true
	}

	conflicts := g.getConflictValues(i, j)

	for _, v := range conflicts {
		if v != 0 {
			g.PossibleValues[i][j][v-1] = false
		}
	}
}

func (g *Grid) Copy() *Grid {
	// TODO There is probably a nicer way to do this
	newgrid := *g
	return &newgrid
}

func (g *Grid) IsSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g.FixedValues[i][j] == 0 {
				return false
			}
		}
	}
	return true
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

func GridFromString(s string) (*Grid, error) {
	var g Grid
	lines := strings.Split(s, "\n")
	if len(lines) != 9 {
		return nil, errors.New("Input string has incorrect number of lines")
	}
	for i, l := range lines {
		for j, c := range l {
			var val int
			if c == '*' {
				val = 0
			} else {
				val = int(c - '0')
			}
			g.FixedValues[i][j] = int8(val)
		}
	}
	g.Update()
	return &g, nil
}
