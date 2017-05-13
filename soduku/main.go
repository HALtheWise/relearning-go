// main.go
package main

import (
	"fmt"
)

var problem = nearlyFull

// Recursive
func solveGrid(g *Grid) (success bool, newgrid *Grid) {
	if g.IsSolved() {
		return true, g
	}

	// Step 1: find the most promising cell to fill
	// Heuristic: the unfilled cell with the fewest possibilities
	// TODO cache this data in the grid and update it dynamically
	var mostPromising struct{ i, j, numOptions int }
	mostPromising.numOptions = 999

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g.FixedValues[i][j] != 0 {
				// This cell is already filled
				continue
			}
			numOptions := 0
			for k := 0; k < 9; k++ {
				if g.PossibleValues[i][j][k] == true {
					numOptions++
				}
			}

			if numOptions < mostPromising.numOptions {
				mostPromising.i = i
				mostPromising.j = j
				mostPromising.numOptions = numOptions
			}
		}
	}

	// Step 2: For each option, consider it by recursing
	i := mostPromising.i
	j := mostPromising.j
	for k := 1; k <= 9; k++ {
		if g.PossibleValues[i][j][k-1] == true {
			// TODO Make sure that grid is getting put on stack instead of heap
			grid := g.Copy()
			grid.Update()
			grid.FixedValues[i][j] = int8(k)

			succ, grid := solveGrid(grid)
			if succ {
				return true, grid
			}
		}
	}

	return false, nil
}

func main() {
	//	fmt.Println("Hello World!")
	fmt.Println(nearlyFull)
	fmt.Println(nearlyFull.PossibleValues[8][8])

	succ, g2 := solveGrid(nearlyFull)
	fmt.Printf("\n\n%v\n%s\n", succ, g2)
}
