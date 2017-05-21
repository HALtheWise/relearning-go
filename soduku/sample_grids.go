// sample_grids
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	blankGrid,
	nearlyFull,
	s04a *Grid
	eulerGrids []*Grid
)

func init() {
	blankGrid, _ = GridFromString(`*********
*********
*********
*********
*********
*********
*********
*********
*********`)
	nearlyFull, _ = GridFromString(`123456789
123456789
123456789
123456789
123456789
123456789
123456789
123456789
19345678*`)
	s04a, _ = GridFromString(`050090000
004800009
000107280
560000137
000000000
173000042
021508000
600003800
000010060`)
}

func init() {
	var err error
	eulerGrids, err = loadEulerGrids("test_data/euler_puzzles.txt")
	if err != nil {
		fmt.Printf("Error loading Euler grids: %v\n", err)
	}
}

func loadEulerGrids(filename string) (grids []*Grid, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	grids = make([]*Grid, 50)
	for i := 0; i < 50; i++ {
		_, _, err = reader.ReadLine()
		if err != nil {
			return
		}
		var griddata []string
		for j := 0; j < 9; j++ {
			line, _ := reader.ReadString('\n')
			griddata = append(griddata, line)
		}
		grids[i], err = GridFromString(strings.TrimSpace(strings.Join(griddata, "")))
		if err != nil {
			return
		}
	}

	return
}
