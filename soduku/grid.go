// grid
package main

import (
	"errors"
	"fmt"
	"strings"
)

type GridCoord struct{ row, col int }
type GridValue int8

type Grid struct {
	FixedValues    [9][9]GridValue
	PossibleValues [9][9][9]bool
}

func (g *Grid) GetFixedValue(coord GridCoord) GridValue {
	return g.FixedValues[coord.row][coord.col]
}

func (g *Grid) GetPossibleValues(coord GridCoord) [9]bool {
	return g.PossibleValues[coord.row][coord.col]
}

func (g *Grid) getConflictValues(coord GridCoord) (conflicts []GridValue) {
	for row := 0; row < 9; row++ {
		// Add the values in my column but not in my row
		if coord.row != row {
			conflicts = append(conflicts, g.FixedValues[row][coord.col])
		}
	}
	for col := 0; col < 9; col++ {
		// Add the values in my row but not in my column
		if coord.col != col {
			conflicts = append(conflicts, g.FixedValues[coord.row][col])
		}
	}

	// Add the values in my box that aren't me
	// TODO this
	root := GridCoord{
		row: (coord.row / 3) * 3,
		col: (coord.col / 3) * 3}

	var test GridCoord
	for test.row = root.row; test.row < root.row+3; test.row++ {
		for test.col = root.col; test.col < root.col+3; test.col++ {
			if test != coord {
				conflicts = append(conflicts, g.GetFixedValue(test))
			}
		}
	}

	return
}

func (g *Grid) Update() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			g.updatePossibilities(GridCoord{i, j})
		}
	}
}

func (g *Grid) updatePossibilities(coord GridCoord) {
	// TODO Be able to do this operation incrementally whenever an element is added to the grid
	// Reset all possiblities to "true"
	for k := 0; k < 9; k++ {
		g.PossibleValues[coord.row][coord.col][k] = true
	}

	conflicts := g.getConflictValues(coord)

	for _, v := range conflicts {
		if v != 0 {
			g.PossibleValues[coord.row][coord.col][v-1] = false
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
			g.FixedValues[i][j] = GridValue(val)
		}
	}
	g.Update()
	return &g, nil
}
