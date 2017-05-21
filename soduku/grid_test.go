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

func TestCanLoadGrid(t *testing.T) {
	_, err := GridFromString(SingleElementGridString)
	if err != nil {
		t.Fatal("Unable to load grid from string %v", SingleElementGridString)
	}
}

func TestNewGridIsEmpty(t *testing.T) {
	grid, _ := GridFromString(EmptyGridString)
	for _, coord := range AllCoords {
		if grid.GetValue(coord) != 0 {
			t.Errorf("Empty grid had non-zero value at coordinate %v", coord)
		}
		for value := GridValue(1); value <= 9; value++ {
			if grid.CanTakeValue(coord, value) != true {
				t.Errorf("Empty grid unable to take value %v at coordinate %v", value, coord)
			}
		}
	}
}

// TestSettingValueUpdatesValue assigns a value of 3 to the top left cell in the grid
// and verifies that when queried that cell returns 3
func TestSettingValueUpdatesValue(t *testing.T) {
	grid, _ := GridFromString(EmptyGridString)

	coord := GridCoord{Row: 0, Col: 0}
	value := GridValue(3)
	grid.SetValue(coord, value)

	if result := grid.GetValue(coord); result != value {
		t.Errorf("After attempting to set value %v, instead got back %v", value, result)
	}
}

func TestSettingValueUpdatesAvailabilitiesCorrectly(t *testing.T) {
	grid, _ := GridFromString(EmptyGridString)

	coord := GridCoord{Row: 0, Col: 0}
	value := GridValue(3)
	grid.SetValue(coord, value)

	tests := []struct {
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

		// Test with different values
		{GridCoord{Row: 0, Col: 6}, GridValue(2), true, "in same row"},
	}

	for _, test := range tests {
		if output := grid.CanTakeValue(test.input, test.value); output != test.output {
			t.Errorf("Cell %v %s as target erroneously returned settability %v for value %v",
				test.input, test.relationship, output, test.value)
		}
	}
}
