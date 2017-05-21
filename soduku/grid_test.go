package main

import (
	"testing"
)

const (
	EmptyGridString = `000000000
000000000
000000000
000000000
000000000
000000000
000000000
000000000
000000000`
	SingleElementGridString = `400000000
000000000
000000000
000000000
000000000
000000000
000000000
000000000
000000000`
)

// TestAllCoords verifies that iterating through AllCoords touches every cell of the grid.
// Warning: this test depends on being able to set contradictory values, which could change
// in the future if the signature of SetValue changes.
func TestAllCoords(t *testing.T) {
	if len(AllCoords) != 81 {
		t.Error("AllCoords does not have length 81")
	}

	grid := NewGrid()
	value := GridValue(1)
	for _, coord := range AllCoords {
		grid.SetValue(coord, value)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if coord := (GridCoord{Row: row, Col: col}); grid.GetValue(coord) != value {
				t.Errorf("Coordinate %v not touched by AllCoords", coord)
			}
		}
	}
}

func verifyGridIsEmpty(grid *Grid) func(t *testing.T) {
	return func(t *testing.T) {
		t.Run("IsUnsolved", func(t *testing.T) {
			if grid.IsSolved() != false {
				t.Error("Empty grid should not report being solved")
			}
		})

		t.Run("IsEmpty", func(t *testing.T) {
			for _, coord := range AllCoords {
				for value := GridValue(1); value <= 9; value++ {
					if grid.CanTakeValue(coord, value) != true {
						t.Errorf("Empty grid unable to take value %v at coordinate %v", value, coord)
					}
				}
			}
		})

		t.Run("AllPossible", func(t *testing.T) {
			for _, coord := range AllCoords {
				for value := GridValue(1); value <= 9; value++ {
					if grid.CanTakeValue(coord, value) != true {
						t.Errorf("Empty grid unable to take value %v at coordinate %v", value, coord)
					}
				}
			}
		})
	}
}

// TestEmptyGrid loads a grid from a string of all 0's and checks
// that it has the expected behavior.
func TestEmptyGrid(t *testing.T) {
	stringGrid, err := GridFromString(EmptyGridString)
	t.Run("LoadFromString", func(t *testing.T) {
		if err != nil {
			t.Fatalf("Unable to load grid from string %v", SingleElementGridString)
		}
	})

	t.Run("LoadedGrid", verifyGridIsEmpty(stringGrid))

	newGrid := NewGrid()

	t.Run("NewGrid", verifyGridIsEmpty(newGrid))
}

// TestSetValue assigns a value of 3 to the top left cell in the grid
// and verifies that when queried that cell returns 3 and
// neighboring cells understand that.
func TestSetValue(t *testing.T) {
	grid, _ := GridFromString(EmptyGridString)

	coord := GridCoord{Row: 0, Col: 0}
	value := GridValue(3)
	grid.SetValue(coord, value)

	// Check that the value we set actually stuck
	t.Run("GetValue", func(t *testing.T) {
		if result := grid.GetValue(coord); result != value {
			t.Errorf("After attempting to set value %v, instead got back %v", value, result)
		}
	})

	// Check that neighboring cells now cannot take that value
	t.Run("CanTakeValue", func(t *testing.T) {
		testPoints := []struct {
			input        GridCoord
			value        GridValue
			output       bool
			relationship string
		}{
			// Test with same value
			{GridCoord{Row: 0, Col: 6}, value, false, "in same row"},
			{GridCoord{Row: 6, Col: 0}, value, false, "in same column"},
			{GridCoord{Row: 1, Col: 2}, value, false, "in same subbox"},
			{GridCoord{Row: 4, Col: 6}, value, true, "elsewhere in grid"},

			// Check that other values weren't barred too
			{GridCoord{Row: 0, Col: 6}, GridValue(2), true, "in same row"},
		}

		for _, test := range testPoints {
			if output := grid.CanTakeValue(test.input, test.value); output != test.output {
				t.Errorf("Cell %v %s as target erroneously returned settability %v for value %v",
					test.input, test.relationship, output, test.value)
			}
		}
	})
}
