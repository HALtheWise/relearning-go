// main.go
package main

import (
	"errors"
	"fmt"
)

// TODO Eliminate all accesses to private fields in this file

// solveGrid is the primary recursive solver for ...
func solveGrid(g Grid) (success bool, newgrid *Grid) {
	if g.IsSolved() {
		heapgrid := g.Clone()
		return true, &heapgrid
	}

	// Step 1: find the most promising cell to fill
	// Heuristic: the unfilled cell with the fewest possibilities
	// TODO cache this data in the grid and update it dynamically
	var mostPromising struct {
		coord      GridCoord
		numOptions int
	}
	mostPromising.numOptions = 999

	for _, coord := range AllCoords {
		if g.GetFixedValue(coord) != 0 {
			// This cell is already filled
			continue
		}
		numOptions := g.getNumOptions(coord)

		if numOptions < mostPromising.numOptions {
			mostPromising.coord = coord
			mostPromising.numOptions = numOptions
		}
	}

	// Step 2: For each option, consider it by recursing
	coord := mostPromising.coord
	for k := GridValue(1); k <= 9; k++ {
		if g.CanTakeValue(coord, k) {
			grid := g.Clone()
			grid.SetValue(coord, k)

			succ, heapgrid := solveGrid(grid)
			if succ {
				return true, heapgrid
			}
		}
	}

	return false, nil
}

func processGrid(grid *Grid) {
	fmt.Println(grid)

	succ, g2 := solveGrid(*grid)
	fmt.Printf("\n\n%v\n%s\n", succ, g2)
}

func solveEuler() (results [][9][9]GridValue, sum int, err error) {
	for i, grid := range eulerGrids {
		succ, newgrid := solveGrid(*grid)
		if !succ {
			return nil, -1, errors.New(fmt.Sprintf("Unable to solve Euler grid #%d", i+1))
		}
		val := int(newgrid.GetFixedValue(GridCoord{0, 0}))*100 +
			int(newgrid.GetFixedValue(GridCoord{0, 1}))*10 +
			int(newgrid.GetFixedValue(GridCoord{0, 2}))
		results = append(results, newgrid.fixedValues)
		sum += val
	}
	return
}

func main() {
	processGrid(s04a)
	//	_, sum, _ := solveEuler()
	//	fmt.Println(sum)
}
