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
	fixedValues    [9][9]GridValue
	possibleValues [9][9][9]bool
}

func (g *Grid) updateConflict(coord GridCoord, other GridValue) {
	g.possibleValues[coord.row][coord.col][other-1] = false
}

func (g *Grid) SetValue(coord GridCoord, value GridValue) {
	g.fixedValues[coord.row][coord.col] = value

	g.updateAllConflicts(coord, value)
}

func (g *Grid) GetFixedValue(coord GridCoord) GridValue {
	return g.fixedValues[coord.row][coord.col]
}

func (g *Grid) GetPossibleValues(coord GridCoord) [9]bool {
	return g.possibleValues[coord.row][coord.col]
}

func (g *Grid) updateAllConflicts(coord GridCoord, value GridValue) {
	for row := 0; row < 9; row++ {
		// Add the values in my column but not in my row
		if coord.row != row {
			g.updateConflict(GridCoord{row: row, col: coord.col}, value)
		}
	}
	for col := 0; col < 9; col++ {
		// Add the values in my row but not in my column
		if coord.col != col {
			g.updateConflict(GridCoord{row: coord.row, col: col}, value)
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
				g.updateConflict(test, value)
			}
		}
	}

	return
}

func (g *Grid) Update() {
	// Reset all possiblities to "true"
	var coord GridCoord
	for coord.row = 0; coord.row < 9; coord.row++ {
		for coord.col = 0; coord.col < 9; coord.col++ {
			for k := 0; k < 9; k++ {
				g.possibleValues[coord.row][coord.col][k] = true
			}
		}
	}

	// Update all possibilities using existing values in the grid
	for coord.row = 0; coord.row < 9; coord.row++ {
		for coord.col = 0; coord.col < 9; coord.col++ {
			value := g.GetFixedValue(coord)

			if value != 0 {
				g.updateAllConflicts(coord, value)
			}
		}
	}
}

func (g *Grid) Clone() *Grid {
	// TODO There is probably a nicer way to do this
	newgrid := *g
	return &newgrid
}

func (g *Grid) IsSolved() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g.fixedValues[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (g *Grid) String() string {
	var lines []string
	for row := 0; row < 9; row++ {
		if row > 0 && row%3 == 0 {
			lines = append(lines, "-----------")
		}
		var s []string
		for col := 0; col < 9; col++ {
			if col > 0 && col%3 == 0 {
				s = append(s, "|")
			}
			val := g.fixedValues[row][col]
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
			g.fixedValues[i][j] = GridValue(val)
		}
	}
	g.Update()
	return &g, nil
}
