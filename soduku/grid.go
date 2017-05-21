// grid
package main

import (
	"errors"
	"fmt"
	"strings"
)

type GridCoord struct{ Row, Col int }
type GridValue int8

var AllCoords []GridCoord

func init() {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			AllCoords = append(AllCoords, GridCoord{row, col})
		}
	}
}

type Grid struct {
	filledCells             int
	fixedValues             [9][9]GridValue
	eliminatedPossibilities [9][9]int8
	possibleValues          [9][9][9]bool
}

func (g *Grid) getNumOptions(coord GridCoord) int {
	return 9 - int(g.eliminatedPossibilities[coord.Row][coord.Col])
}

func (g *Grid) CanTakeValue(coord GridCoord, value GridValue) bool {
	return g.possibleValues[coord.Row][coord.Col][value-1]
}

func (g *Grid) updateConflict(coord GridCoord, other GridValue) {
	wasPossible := g.possibleValues[coord.Row][coord.Col][other-1]
	if wasPossible {
		g.possibleValues[coord.Row][coord.Col][other-1] = false
		g.eliminatedPossibilities[coord.Row][coord.Col]++
	}
}
func (g *Grid) SetValue(coord GridCoord, value GridValue) {
	g.filledCells++
	g.fixedValues[coord.Row][coord.Col] = value

	g.updateAllConflicts(coord, value)
}

func (g *Grid) GetValue(coord GridCoord) GridValue {
	return g.fixedValues[coord.Row][coord.Col]
}

func (g *Grid) updateAllConflicts(coord GridCoord, value GridValue) {
	// Consider the values in my column
	test := coord
	for test.Row = 0; test.Row < 9; test.Row++ {
		g.updateConflict(test, value)
	}

	// Consider the values in my row
	test = coord
	for test.Col = 0; test.Col < 9; test.Col++ {
		g.updateConflict(test, value)
	}

	// Consider the values in my 3x3 box
	root := GridCoord{
		Row: (coord.Row / 3) * 3,
		Col: (coord.Col / 3) * 3}

	for test.Row = root.Row; test.Row < root.Row+3; test.Row++ {
		for test.Col = root.Col; test.Col < root.Col+3; test.Col++ {
			g.updateConflict(test, value)
		}
	}
}

func (g *Grid) Update() {
	// Reset all possiblities to "true"
	for _, coord := range AllCoords {
		for k := 0; k < 9; k++ {
			g.possibleValues[coord.Row][coord.Col][k] = true
		}
	}

	g.filledCells = 0

	// Update all possibilities using existing values in the grid
	for _, coord := range AllCoords {
		value := g.GetValue(coord)

		if value != 0 {
			g.filledCells++
			g.updateAllConflicts(coord, value)
		}
	}
}

func (g *Grid) Clone() (newgrid Grid) {
	// TODO There is probably a nicer way to do this that is stack-compatible
	newgrid = *g
	return
}

func (g *Grid) IsSolved() bool {
	return g.filledCells == 81
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
			val := g.GetValue(GridCoord{Row: row, Col: col})
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
